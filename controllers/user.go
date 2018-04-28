package controllers

import (
	"github.com/astaxie/beego/context"
	"encoding/json"
	"github.com/vibakar/golang-movieapp/models/database"
	"golang.org/x/crypto/bcrypt"
	"github.com/satori/go.uuid"
	"github.com/astaxie/beego"
)

type movieData struct {
	Id int `json:"id"`
	Title string `json:"title"`
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
var sessionDB = map[string]string{}

func AddMovie(ctx *context.Context){
	db, err := database.ConnectDB()
	if err == nil {
		defer db.Close()
		var movie movieData
		data := ctx.Input.RequestBody
		json.Unmarshal(data, &movie)
		insert, err := db.Prepare("INSERT INTO favourites(id, title, votes, rating, poster) VALUES(?,?,?,?,?)")
		if err == nil {
			_, err := insert.Exec(movie.Id, movie.Title, movie.Votes, movie.Rating, movie.Poster)
			if err == nil {
				beego.Info("Moive added to fav succeess")
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
}

func GetFavMovies(ctx *context.Context){
	db, err := database.ConnectDB();
	var moviesList = make([]interface{}, 0)
	if err == nil {
		defer db.Close()
		rows, err := db.Query("SELECT * FROM favourites")
		if err == nil {
			for rows.Next(){
				var movie movieData
				rows.Scan(&movie.Id, &movie.Title, &movie.Votes, &movie.Rating, &movie.Poster)
				moviesList = append(moviesList, movie)
			}
			beego.Info("favourite movie fetched successfully", moviesList)
			var response, _ = json.Marshal(moviesList)
			ctx.Output.Body(response)
		} else {
			beego.Error("Get fav movie query fails")
			ctx.Output.Status = 503
			ctx.Output.Body([]byte(`{"errMsg": "Service Unavailable, Try Later"}`))
		}
	} else {
		beego.Error("DB connection failed during get favourite movies")
		ctx.Output.Status = 503
		ctx.Output.Body([]byte(`{"errMsg": "Service Unavailable, Try Later"}`))
	}
}

func DeleteMovie(ctx *context.Context){
	db, err := database.ConnectDB()
	if err == nil {
		var movieId = ctx.Input.Param(":movieId")
		del, err := db.Prepare("DELETE FROM favourites WHERE id=?")
		if err == nil {
			del.Exec(movieId)
			beego.Info("Removing movie from favourites success")
			ctx.Output.Body([]byte(`{"response": "Movie removed from favourites"}`))
		} else {
			beego.Error("Delete favourite movie query fails")
			ctx.Output.Status = 500
			ctx.Output.Body([]byte(`{"errMsg": "Failed to remove movie from favourites"}`))
		}
	} else {
		beego.Error("DB connection failed during delete movie from favourites")
		ctx.Output.Status = 503
		ctx.Output.Body([]byte(`{"errMsg": "Service Unavailable, Try Later"}`))
	}
}

func Signup(ctx *context.Context){
	db, err := database.ConnectDB()
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
				ctx.SetCookie("U_SESSION_ID", uid.String())
				sessionDB[uid.String()] = signupData.Email
				beego.Info("User signup success with email ", signupData.Email)
				ctx.Output.Status = 201
				ctx.Output.Body([]byte(`{"response": "Account created successfully"}`))
			} else {
				beego.Warn("user signup failed because of using already available email", signupData.Email)
				ctx.Output.Status = 409
				ctx.Output.Body([]byte(`{"errMsg": "Email already exists"}`))
			}
		} else {
			beego.Error("user signup query fails")
			ctx.Output.Status = 503
			ctx.Output.Body([]byte(`{"errMsg": "Service Unavailable, Try Later"}`))
		}
	} else {
		beego.Error("DB connection failed during user signup")
		ctx.Output.Status = 503
		ctx.Output.Body([]byte(`{"errMsg": "Service Unavailable, Try Later"}`))
	}
}

func Login(ctx *context.Context)  {
	db, err := database.ConnectDB()
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
				ctx.SetCookie("U_SESSION_ID", uid.String())
				sessionDB[uid.String()] = loginData.Email
				beego.Info("User logged in successfully with Email", loginData.Email)
				ctx.Output.Status = 200
				ctx.Output.Body([]byte(`{"response": "Login success"}`))
			} else {
				beego.Error("User login failed", loginData.Email)
				ctx.Output.Status = 401
				ctx.Output.Body([]byte(`{"errMsg": "Email or Password incorrect"}`))
			}
		} else {
			beego.Error("User login query failed")
			ctx.Output.Status = 503
			ctx.Output.Body([]byte(`{"errMsg": "Service Unavailable, Try Later"}`))
		}
	} else {
		beego.Error("DB connection failed in user login")
		ctx.Output.Status = 503
		ctx.Output.Body([]byte(`{"errMsg": "Service Unavailable, Try Later"}`))
	}
}

func Logout(ctx *context.Context){
	uid := ctx.GetCookie("U_SESSION_ID")
	delete(sessionDB, uid)
	beego.Info("User logged out and session cleared ", uid)
	ctx.Output.Body([]byte(`{"response":"Logged out successfully"}`))
}

func GetUsername(ctx *context.Context){
	uid := ctx.GetCookie("U_SESSION_ID")
	if len(uid) > 0 {
		email := sessionDB[uid]
		if len(email) > 0 {
			db, err := database.ConnectDB()
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
					ctx.Output.Body([]byte(`{"errMsg": "Service Unavailable, Try Later"}`))
				}
			} else {
				beego.Error("DB connection failed while getting username")
				ctx.Output.Status = 503
				ctx.Output.Body([]byte(`{"errMsg": "Service Unavailable, Try Later"}`))
			}
		} else {
			beego.Error("No email found for the received cookie")
			ctx.Output.Status = 401
			ctx.Output.Body([]byte(`{"errMsg": "Unauthorised user"}`))
		}
	} else {
		beego.Error("No cookie recieved to get username")
		ctx.Output.Status = 401
		ctx.Output.Body([]byte(`{"errMsg": "Unauthorised user"}`))
	}
}