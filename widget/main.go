package main

import (
	_ "widget/routers"
	"github.com/astaxie/beego"
	"os"
)

func main() {
	beego.SetStaticPath(os.Getenv("WIDGET_URL_PREFIX") + "static","static")
	beego.Run()
}

