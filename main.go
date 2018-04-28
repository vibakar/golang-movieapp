package main

import (
	_ "github.com/vibakar/golang-movieapp/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetLogger("file", `{"filename":"logs/app.log"}`)
	beego.Run()
}

