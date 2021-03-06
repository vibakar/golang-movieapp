package controllers

import (
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego"
	"net/url"
	"time"
)

var apikey = beego.AppConfig.String("TMDBApiKey")
var reqTimeOut = 30 * time.Second

func GetNowPlayingMovies(ctx *context.Context) {
	req := httplib.Get("https://api.themoviedb.org/3/movie/now_playing?api_key="+apikey+"&language=en-US&page=1").SetTimeout(reqTimeOut, 30 * time.Second)
	data, err := req.Bytes()
	if err == nil {
		beego.Info("Now playing movies successfully fetched from TMDB")
		ctx.Output.Body(data)
	} else {
		beego.Error("Failed to fetch Now playing movies from TMDB")
		ctx.Output.Status = 500
		ctx.Output.Body([]byte(`{"errMsg": "Internal Server Error", "code": 500}`))
	}
}

func GetTopRatedMovies(ctx *context.Context) {
	req := httplib.Get("https://api.themoviedb.org/3/movie/top_rated?api_key="+apikey+"&language=en-US&page=1").SetTimeout(reqTimeOut, 30 * time.Second)
	data, err := req.Bytes()
	if err == nil {
		beego.Info("Top Rated movies successfully fetched from TMDB")
		ctx.Output.Body(data)
	} else {
		beego.Error("Failed to fetch Top Rated movies from TMDB")
		ctx.Output.Status = 500
		ctx.Output.Body([]byte(`{"errMsg": "Internal Server Error", "code": 500}`))
	}
}

func GetUpcomingMovies(ctx *context.Context) {
	req := httplib.Get("https://api.themoviedb.org/3/movie/upcoming?api_key="+apikey+"&language=en-US&page=1").SetTimeout(reqTimeOut, 30 * time.Second)
	data, err := req.Bytes()
	if err == nil {
		beego.Info("Upcoming movies successfully fetched from TMDB")
		ctx.Output.Body(data)
	} else {
		beego.Error("Failed to fetch upcoming movies from TMDB")
		ctx.Output.Status = 500
		ctx.Output.Body([]byte(`{"errMsg": "Internal Server Error", "code": 500}`))
	}
}

func GetSearchedMovies(ctx *context.Context) {
	var movie string
	ctx.Input.Bind(&movie, "movie")
	beego.Info("searched movie name", movie)
	req := httplib.Get("https://api.themoviedb.org/3/search/movie?api_key="+apikey+"&language=en-US&query="+url.QueryEscape(movie)).SetTimeout(reqTimeOut, 30 * time.Second)
	data, err := req.Bytes()
	if err == nil {
		beego.Info("Successfully fetched serached movies from TMDB")
		ctx.Output.Body(data)
	} else {
		beego.Error("Failed to get searched movies from TMDB")
		ctx.Output.Status = 500
		ctx.Output.Body([]byte(`{"errMsg": "Internal Server Error", "code": 500}`))
	}
}

func GetMovieDetail(ctx *context.Context) {
	movieId := ctx.Input.Param(":movieId")
	beego.Info("Received movie Id", movieId)
	req := httplib.Get("https://api.themoviedb.org/3/movie/"+movieId+"?api_key="+apikey+"&language=en-US").SetTimeout(reqTimeOut, 30 * time.Second)
	data, err := req.Bytes()
	if err == nil {
		beego.Info("Successfully fetched movie detail from TMDB")
		ctx.Output.Body(data)
	} else {
		beego.Error("Failed to get movies detail from TMDB", movieId)
		ctx.Output.Status = 500
		ctx.Output.Body([]byte(`{"errMsg": "Internal Server Error", "code": 500}`))
	}
}

func GetSimilarMovies(ctx *context.Context){
	movieId := ctx.Input.Param(":movieId")
	req := httplib.Get("https://api.themoviedb.org/3/movie/"+movieId+"/similar?api_key="+apikey+"&language=en-US&page=1").SetTimeout(reqTimeOut, 30 * time.Second)
	data, err := req.Bytes()
	if err == nil {
		beego.Info("Successfully fetched similar movies from TMDB")
		ctx.Output.Body(data)
	} else {
		beego.Error("Failed to get similar movies from TMDB")
		ctx.Output.Status = 500
		ctx.Output.Body([]byte(`{"errMsg": "Internal Server Error", "code": 500}`))
	}
}