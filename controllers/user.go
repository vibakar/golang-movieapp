package controllers

import (
	"github.com/astaxie/beego/context"
	"encoding/json"
	"github.com/vibakar/golang-movieapp/models/database"
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
				ctx.Output.Status = 201
				ctx.Output.Body([]byte(`{"response": "Movie added to favourites"}`))
			} else {
				ctx.Output.Status = 409
				ctx.Output.Body([]byte(`{"errMsg": "Movie already added to favourites"}`))
			}
		} else {
			ctx.Output.Status = 503
			ctx.Output.Body([]byte(`{"errMsg": "Failed to add movie to favourites"}`))
		}
	} else {
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
			var response, _ = json.Marshal(moviesList)
			ctx.Output.Body(response)
		} else {
			ctx.Output.Status = 500
			ctx.Output.Body([]byte(`{"errMsg": "Failed to fetch fav movies"}`))
		}
	} else {
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
			ctx.Output.Body([]byte(`{"response": "Movie removed from favourites"}`))
		} else {
			ctx.Output.Status = 500
			ctx.Output.Body([]byte(`{"errMsg": "Failed to remove movie from favourites"}`))
		}
	} else {
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
		insert, err := db.Prepare("INSERT INTO user(username, email, password) VALUES(?,?,?)")
		if err == nil {
			_, err := insert.Exec(signupData.Username, signupData.Email, signupData.Password)
			if err == nil {
				ctx.Output.Status = 201
				ctx.Output.Body([]byte(`{"response": "Account created successfully"}`))
			} else {
				ctx.Output.Status = 409
				ctx.Output.Body([]byte(`{"errMsg": "Email already exists"}`))
			}
		} else {
			ctx.Output.Status = 503
			ctx.Output.Body([]byte(`{"errMsg": "Service Unavailable, Try Later"}`))
		}
	} else {
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
			if dbPassword == loginData.Password {
				ctx.Output.Status = 200
				ctx.Output.Body([]byte(`{"response": "Login success"}`))
			} else {
				ctx.Output.Status = 401
				ctx.Output.Body([]byte(`{"errMsg": "Email or Password incorrect"}`))
			}
		} else {
			ctx.Output.Status = 503
			ctx.Output.Body([]byte(`{"errMsg": "Service Unavailable, Try Later"}`))
		}
	} else {
		ctx.Output.Status = 503
		ctx.Output.Body([]byte(`{"errMsg": "Service Unavailable, Try Later"}`))
	}
}