package routers

import (
	"github.com/vibakar/golang-movieapp/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Get("/v1/movie/nowPlaying", controllers.GetNowPlayingMovies)
    beego.Get("/v1/movie/topRated", controllers.GetTopRatedMovies)
    beego.Get("/v1/movie/upcoming", controllers.GetUpcomingMovies)
    beego.Get("/v1/movie/search?:movie", controllers.GetSearchedMovies)
    beego.Get("/v1/movie/detail/:movieId", controllers.GetMovieDetail)
    beego.Get("/v1/movie/similar/:movieId", controllers.GetSimilarMovies)

    beego.Post("/v1/user/signUp", controllers.SignUp)
    beego.Post("/v1/user/login", controllers.Login)
    beego.Get("/v1/user/logout", controllers.Logout)
    beego.Post("/v1/user/addMovie", controllers.AddMovie)
    beego.Get("/v1/user/favMovies", controllers.GetFavMovies)
    beego.Delete("/v1/user/delMovie/:movieId", controllers.DeleteMovie)
    beego.Get("/v1/user/username", controllers.GetUsername)
    beego.Post("/v1/user/verifyEmail", controllers.VerifyEmail)
    beego.Get("/v1/user/resendCode", controllers.ResendCode)
}