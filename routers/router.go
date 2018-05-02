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

    beego.Post("/v1/user/signup", controllers.SignUp)
    beego.Post("/v1/user/login", controllers.Login)
    beego.Get("/v1/user/logout", controllers.Logout)
    beego.Post("/v1/user/addmovie", controllers.AddMovie)
    beego.Get("/v1/user/favmovies", controllers.GetFavMovies)
    beego.Delete("/v1/user/delmovie/:movieId", controllers.DeleteMovie)
    beego.Get("/v1/user/username", controllers.GetUsername)
    beego.Post("/v1/user/verifyEmail", controllers.VerifyEmail)
}