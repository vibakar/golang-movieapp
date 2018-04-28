package controllers

import (
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/httplib"
	"encoding/json"
	"github.com/astaxie/beego"
)

var apikey = beego.AppConfig.String("tmdb_apikey")

type ErrorDetail struct {
	Status int `json:"status"`
	ErrorMessage string `json:"errorMessage"`
}

func GetNowPlayingMovies(ctx *context.Context) {
	req := httplib.Get("https://api.themoviedb.org/3/movie/now_playing?api_key="+apikey+"&language=en-US&page=1")
	data, err := req.Bytes()
	if err == nil {
		beego.Info("Now playing movies successfully fetched from TMDB")
		ctx.Output.Body(data)
	} else {
		beego.Error("Failed to fetch Now playing movies from TMDB")
		ctx.Output.Status = 500
		var error = ErrorDetail{500, "Internal Server Error"}
		jsonErr, _ := json.Marshal(error)
		ctx.Output.Body(jsonErr)
	}
}

func GetTopRatedMovies(ctx *context.Context) {
	req := httplib.Get("https://api.themoviedb.org/3/movie/top_rated?api_key="+apikey+"&language=en-US&page=1")
	data, err := req.Bytes()
	if err == nil {
		beego.Info("Top Rated movies successfully fetched from TMDB")
		ctx.Output.Body(data)
	} else {
		beego.Error("Failed to fetch Top Rated movies from TMDB")
		ctx.Output.Status = 500
		var error = ErrorDetail{500, "Internal Server Error"}
		jsonErr, _ := json.Marshal(error)
		ctx.Output.Body(jsonErr)
	}
}

func GetUpcomingMovies(ctx *context.Context) {
	req := httplib.Get("https://api.themoviedb.org/3/movie/upcoming?api_key="+apikey+"&language=en-US&page=1")
	data, err := req.Bytes()
	if err == nil {
		beego.Info("Upcoming movies successfully fetched from TMDB")
		ctx.Output.Body(data)
	} else {
		beego.Error("Failed to fetch upcoming movies from TMDB")
		ctx.Output.Status = 500
		var error = ErrorDetail{500, "Internal Server Error"}
		jsonErr, _ := json.Marshal(error)
		ctx.Output.Body(jsonErr)
	}
}

func GetSearchedMovies(ctx *context.Context) {
	var movie string
	ctx.Input.Bind(&movie, "movie")
	req := httplib.Get("https://api.themoviedb.org/3/search/movie?api_key="+apikey+"&language=en-US&query="+movie)
	data, err := req.Bytes()
	if err == nil {
		beego.Info("Successfully fetched serached movies from TMDB")
		ctx.Output.Body(data)
	} else {
		beego.Error("Failed to get searched movies from TMDB")
		ctx.Output.Status = 500;
		var error = ErrorDetail{500, "Internal Server Error"}
		jsonErr, _ := json.Marshal(error)
		ctx.Output.Body(jsonErr)
	}
}

func GetSimilarMovies(ctx *context.Context){
	var movieId = ctx.Input.Param(":movieId")
	req := httplib.Get("https://api.themoviedb.org/3/movie/"+movieId+"/similar?api_key="+apikey+"&language=en-US&page=1")
	data, err := req.Bytes()
	if err == nil {
		beego.Info("Successfully fetched similar movies from TMDB")
		ctx.Output.Body(data)
	} else {
		beego.Error("Failed to get similar movies from TMDB")
		ctx.Output.Status = 500;
		var error = ErrorDetail{500, "Internal Server Error"}
		jsonErr, _ := json.Marshal(error)
		ctx.Output.Body(jsonErr)
	}
}