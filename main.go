package main

import (
	_ "beego_project/routers"
	"github.com/astaxie/beego"
)

func main() {
	//beego.SetViewsPath("front")
	beego.Run()
}