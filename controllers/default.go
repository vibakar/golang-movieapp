package controllers

import (
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/httplib"
	"encoding/json"
	"github.com/astaxie/beego"
)

var apikey string = beego.AppConfig.String("tmdb_apikey")

type ErrorDetail struct {
	Status int `json:"status"`
	ErrorMessage string `json:"errorMessage"`
}

func GetNowPlayingMovies(ctx *context.Context) {
	req := httplib.Get("https://api.themoviedb.org/3/movie/now_playing?api_key="+apikey+"&language=en-US&page=1")
	data, err := req.Bytes()
	if err == nil {
		ctx.Output.Body(data)
	} else {
		ctx.Output.Status = 500;
		var error = ErrorDetail{500, "Internal Server Error"}
		jsonErr, _ := json.Marshal(error);
		ctx.Output.Body(jsonErr)
	}
}

func GetTopRatedMovies(ctx *context.Context) {
	req := httplib.Get("https://api.themoviedb.org/3/movie/top_rated?api_key="+apikey+"&language=en-US&page=1")
	data, err := req.Bytes()
	if err == nil {
		ctx.Output.Body(data)
	} else {
		ctx.Output.Status = 500;
		var error = ErrorDetail{500, "Internal Server Error"}
		jsonErr, _ := json.Marshal(error);
		ctx.Output.Body(jsonErr)
	}
}

func GetUpcomingMovies(ctx *context.Context) {
	req := httplib.Get("https://api.themoviedb.org/3/movie/upcoming?api_key="+apikey+"&language=en-US&page=1")
	data, err := req.Bytes()
	if err == nil {
		ctx.Output.Body(data)
	} else {
		ctx.Output.Status = 500;
		var error = ErrorDetail{500, "Internal Server Error"}
		jsonErr, _ := json.Marshal(error);
		ctx.Output.Body(jsonErr)
	}
}

func GetSearchedMovies(ctx *context.Context) {
	var movie string;
	ctx.Input.Bind(&movie, "movie")
	req := httplib.Get("https://api.themoviedb.org/3/search/movie?api_key="+apikey+"&language=en-US&query="+movie)
	data, err := req.Bytes()
	if err == nil {
		ctx.Output.Body(data)
	} else {
		ctx.Output.Status = 500;
		var error = ErrorDetail{500, "Internal Server Error"}
		jsonErr, _ := json.Marshal(error);
		ctx.Output.Body(jsonErr)
	}
}