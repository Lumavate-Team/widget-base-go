package main

import (
  "github.com/Lumavate-Team/lumavate-go-common/properties/component_data"
	_ "widget/routers"
	"github.com/astaxie/beego"
  "html/template"
	"os"
)

func ComponentHtml(in component_data.ComponentData) (out template.HTML){
    out = template.HTML(in.GetHtml())
    return
}

func StaticPath() (out string) {
	out = os.Getenv("WIDGET_URL_PREFIX") + os.Getenv("PUBLIC_KEY") + "/static"
	return
}

func main() {
	beego.SetStaticPath(os.Getenv("WIDGET_URL_PREFIX") + os.Getenv("PUBLIC_KEY") + "/static","static")
	beego.SetStaticPath(os.Getenv("WIDGET_URL_PREFIX") + os.Getenv("PUBLIC_KEY") +  "/lc","/lumavate-components/dist")
	beego.AddFuncMap("componentHtml", ComponentHtml)
	beego.AddFuncMap("staticPath", StaticPath)
	beego.Run()
}

