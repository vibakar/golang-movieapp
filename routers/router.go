package routers

import (
	"github.com/vibakar/golang-movieapp/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Get("/v1/movie/nowplaying", controllers.GetNowPlayingMovies)
}