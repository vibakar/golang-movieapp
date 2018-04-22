package routers

import (
	"github.com/vibakar/golang-movieapp/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Get("/v1/movie/nowplaying", controllers.GetNowPlayingMovies)
    beego.Get("/v1/movie/toprated", controllers.GetTopRatedMovies)
    beego.Get("/v1/movie/upcoming", controllers.GetUpcomingMovies)
    beego.Get("/v1/movie/search?:movie", controllers.GetSearchedMovies)
    beego.Get("/v1/movie/similar/:movieId", controllers.GetSimilarMovies)
    beego.Post("/v1/users/addmovie", controllers.AddMovie)
}