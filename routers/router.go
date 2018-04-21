package routers

import (
	"github.com/vibakar/movie-webapp/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
