package controllers

import (
	"github.com/astaxie/beego/context"
	"encoding/json"
	"github.com/vibakar/golang-movieapp/models/database"
)

type movieData struct {
	Id int
	Title string
	Votes int
	Rating float64
	Poster string
}

func AddMovie(ctx *context.Context){
	var db, err = database.ConnectDB()
	if err == nil {
		defer db.Close()
		var movie movieData
		data := ctx.Input.RequestBody
		json.Unmarshal(data, &movie)
		insert, err := db.Prepare("INSERT INTO user(id, title, votes, rating, poster) VALUES(?,?,?,?,?)")
		if err == nil {
			_, err := insert.Exec(movie.Id, movie.Title, movie.Votes, movie.Rating, movie.Poster)
			if err == nil {
				ctx.Output.Status = 201
				ctx.Output.Body([]byte(`{"response": "Movie added to favourites"}`))
			} else {
				ctx.Output.Status = 409
				ctx.Output.Body([]byte(`{"response": "Movie already added to favourites"}`))
			}
		} else {
			ctx.Output.Status = 500
			ctx.Output.Body([]byte(`{"errMsg": "Failed to add movie to favourites"}`))
		}
	} else {
		ctx.Output.Status = 503
		ctx.Output.Body([]byte(`{"errMsg": "Service Unavailable"}`))
	}

}