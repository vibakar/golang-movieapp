package controllers

import (
	"github.com/astaxie/beego/context"
	"encoding/json"
	"github.com/vibakar/golang-movieapp/models/dbConnection"
	"golang.org/x/crypto/bcrypt"
	"github.com/satori/go.uuid"
	"github.com/astaxie/beego"
	"github.com/vibakar/golang-movieapp/models/etcd"
)

type movieData struct {
	id int
	email string
	Title string `json:"title"`
	MovieId int `json:"id"`
	Votes int `json:"votes"`
	Rating float64 `json:"rating"`
	Poster string `json:"poster"`
}

type signupData struct {
	Username string
	Email string
	Password string
}

type loginData struct {
	Email string
	Password string
}

type username struct {
	Username string `json:"username"`
}

var cookieMaxAge, _ = beego.AppConfig.Int("cookieMaxAge")

func AddMovie(ctx *context.Context){
	cookie := ctx.GetCookie("U_SESSION_ID")
	resp, err := etcd.Get(cookie)
	if err == nil {
		email := resp.Node.Value
		db, err := dbConnection.ConnectToMysql()
		if err == nil {
			defer db.Close()
			var movie movieData
			data := ctx.Input.RequestBody
			json.Unmarshal(data, &movie)
			insert, err := db.Prepare("INSERT INTO favourites(email, movieId, title, votes, rating, poster) VALUES(?,?,?,?,?,?)")
			if err == nil {
				_, err := insert.Exec(email, movie.MovieId, movie.Title, movie.Votes, movie.Rating, movie.Poster)
				if err == nil {
					beego.Info("Movie added to fav success")
					ctx.Output.Status = 201
					ctx.Output.Body([]byte(`{"response": "Movie added to favourites"}`))
				} else {
					beego.Info("Add movie to fav fails, because movie already exists")
					ctx.Output.Status = 409
					ctx.Output.Body([]byte(`{"errMsg": "Movie already added to favourites"}`))
				}
			} else {
				beego.Error("Add movie to fav query failed")
				ctx.Output.Status = 503
				ctx.Output.Body([]byte(`{"errMsg": "Failed to add movie to favourites"}`))
			}
		} else {
			beego.Error("DB connection failed during add movies to fav")
			ctx.Output.Status = 503
			ctx.Output.Body([]byte(`{"errMsg": "Service Unavailable, Try Later"}`))
		}
	} else {
		beego.Error("Tried to add movie to fav without login")
		ctx.Output.Status = 403
		ctx.Output.Body([]byte(`{"errMsg": "Please login first", "code": 403}`))
	}
}

func GetFavMovies(ctx *context.Context){
	cookie := ctx.GetCookie("U_SESSION_ID")
	resp, err := etcd.Get(cookie)
	if err == nil {
		email := resp.Node.Value
		db, err := dbConnection.ConnectToMysql()
		var moviesList = make([]interface{}, 0)
		if err == nil {
			defer db.Close()
			rows, err := db.Query("SELECT * FROM favourites WHERE email = ?", email)
			if err == nil {
				for rows.Next(){
					var movie movieData
					rows.Scan(&movie.id, &movie.email, &movie.MovieId, &movie.Title, &movie.Votes, &movie.Rating, &movie.Poster)
					moviesList = append(moviesList, movie)
				}
				beego.Info("favourite movie fetched successfully", moviesList)
				var response, _ = json.Marshal(moviesList)
				ctx.Output.Body(response)
			} else {
				beego.Error("Get fav movie query fails")
				ctx.Output.Status = 503
				ctx.Output.Body([]byte(`{"errMsg": "Service Unavailable, Try Later", "code": 503}`))
			}
		} else {
			beego.Error("DB connection failed during get favourite movies")
			ctx.Output.Status = 503
			ctx.Output.Body([]byte(`{"errMsg": "Service Unavailable, Try Later", "code": 503}`))
		}
	} else {
		beego.Error("Tried to get fav movies without login")
		ctx.Output.Status = 403
		ctx.Output.Body([]byte(`{"errMsg": "Access Forbidden", "code": 403}`))
	}
}

func DeleteMovie(ctx *context.Context){
	cookie := ctx.GetCookie("U_SESSION_ID")
	resp, err := etcd.Get(cookie)
	if err == nil {
		email := resp.Node.Value
		db, err := dbConnection.ConnectToMysql()
		if err == nil {
			var movieId = ctx.Input.Param(":movieId")
			del, err := db.Prepare("DELETE FROM favourites WHERE movieId=? AND email=?")
			if err == nil {
				del.Exec(movieId, email)
				beego.Info("Removing movie from favourites success")
				ctx.Output.Body([]byte(`{"response": "Movie removed from favourites"}`))
			} else {
				beego.Error("Delete favourite movie query fails")
				ctx.Output.Status = 500
				ctx.Output.Body([]byte(`{"errMsg": "Failed to remove movie from favourites", "code": 500}`))
			}
		} else {
			beego.Error("DB connection failed during delete movie from favourites")
			ctx.Output.Status = 503
			ctx.Output.Body([]byte(`{"errMsg": "Service Unavailable, Try Later", "code": 503}`))
		}
	} else {
		beego.Error("Tried to delete fav movies without login")
		ctx.Output.Status = 403
		ctx.Output.Body([]byte(`{"errMsg": "Access Forbidden", "code": 403}`))
	}
}

func SignUp(ctx *context.Context){
	db, err := dbConnection.ConnectToMysql()
	if err == nil {
		defer db.Close()
		var signupData signupData
		reqBody := ctx.Input.RequestBody
		json.Unmarshal(reqBody, &signupData)
		hash, _ := bcrypt.GenerateFromPassword([]byte(signupData.Password), bcrypt.MinCost)
		insert, err := db.Prepare("INSERT INTO user(username, email, password) VALUES(?,?,?)")
		if err == nil {
			_, err := insert.Exec(signupData.Username, signupData.Email, hash)
			if err == nil {
				uid, _ := uuid.NewV4()
				ctx.SetCookie("U_SESSION_ID", uid.String(), cookieMaxAge)
				_, err := etcd.Set(uid.String(), signupData.Email)
				if err == nil {
					beego.Info("User signup success with email ", signupData.Email)
					ctx.Output.Status = 201
					ctx.Output.Body([]byte(`{"response": "Account created successfully", "code": 201}`))
				} else {
					beego.Error("Failed to set key in ETCD during signup")
					ctx.Output.Status = 503
					ctx.Output.Body([]byte(`{"errMsg": "Service Unavailable, Try Later", "code": 503}`))
				}
			} else {
				beego.Warn("user signup failed because of using already available email", signupData.Email)
				ctx.Output.Status = 409
				ctx.Output.Body([]byte(`{"errMsg": "Email already exists", "code": 409}`))
			}
		} else {
			beego.Error("user signup query fails")
			ctx.Output.Status = 503
			ctx.Output.Body([]byte(`{"errMsg": "Service Unavailable, Try Later", "code": 503}`))
		}
	} else {
		beego.Error("DB connection failed during user signup")
		ctx.Output.Status = 503
		ctx.Output.Body([]byte(`{"errMsg": "Service Unavailable, Try Later", "code": 503}`))
	}
}

func Login(ctx *context.Context)  {
	db, err := dbConnection.ConnectToMysql()
	if err == nil {
		defer db.Close()
		reqData := ctx.Input.RequestBody
		var loginData loginData
		json.Unmarshal(reqData, &loginData)
		var dbPassword string
		err := db.QueryRow("SELECT password FROM user WHERE email = ?", loginData.Email).Scan(&dbPassword)
		if err == nil {
			err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(loginData.Password))
			if err == nil {
				uid, _ := uuid.NewV4()
				ctx.SetCookie("U_SESSION_ID", uid.String(), cookieMaxAge)
				_, err := etcd.Set(uid.String(), loginData.Email)
				if err == nil {
					beego.Info("User logged in successfully with Email", loginData.Email)
					ctx.Output.Status = 200
					ctx.Output.Body([]byte(`{"response": "Login success"}`))
				} else {
					beego.Error("Failed to set key to ETCD during login")
					ctx.Output.Status = 503
					ctx.Output.Body([]byte(`{"errMsg": "Service Unavailable, Try Later", "code": 503}`))
				}
			} else {
				beego.Error("User login failed", loginData.Email)
				ctx.Output.Status = 401
				ctx.Output.Body([]byte(`{"errMsg": "Email or Password incorrect", "code": 401}`))
			}
		} else {
			beego.Error("User login query failed")
			ctx.Output.Status = 503
			ctx.Output.Body([]byte(`{"errMsg": "Service Unavailable, Try Later", "code": 503}`))
		}
	} else {
		beego.Error("DB connection failed in user login")
		ctx.Output.Status = 503
		ctx.Output.Body([]byte(`{"errMsg": "Service Unavailable, Try Later", "code": 503}`))
	}
}

func Logout(ctx *context.Context){
	cookie := ctx.GetCookie("U_SESSION_ID")
	_, err := etcd.Delete(cookie)
	if err == nil {
		beego.Info("User logged out and session cleared ", cookie)
		ctx.Output.Body([]byte(`{"response":"Logged out successfully"}`))
	} else {
		beego.Error("Failed to delete key in ETCD during logout")
		ctx.Output.Status = 503
		ctx.Output.Body([]byte(`{"errMsg": "Service Unavailable, Try Later", "code": 503}`))
	}
}

func GetUsername(ctx *context.Context){
	cookie := ctx.GetCookie("U_SESSION_ID")
	resp, err := etcd.Get(cookie)
	if err == nil {
		email := resp.Node.Value
		db, err := dbConnection.ConnectToMysql()
		defer db.Close()
		var user username
		if err == nil {
			err := db.QueryRow("SELECT username FROM user WHERE email = ?", email).Scan(&user.Username)
			if err == nil {
				beego.Info("username found for the received cookie")
				ctx.Output.Status = 200
				res,_ := json.Marshal(user)
				ctx.Output.Body(res)
			} else {
				beego.Error("Username not found for the received cookie")
				ctx.Output.Status = 503
				ctx.Output.Body([]byte(`{"errMsg": "Service Unavailable, Try Later", "code": 503}`))
			}
		} else {
			beego.Error("DB connection failed while getting username")
			ctx.Output.Status = 503
			ctx.Output.Body([]byte(`{"errMsg": "Service Unavailable, Try Later", "code": 503}`))
		}
	} else {
		beego.Error("No user found")
		ctx.Output.Status = 403
		ctx.Output.Body([]byte(`{"errMsg": "Unauthorised user", "code": 401}`))
	}
}