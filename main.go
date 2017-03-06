package main

import (
	_ "wx_server_go/docs"
	_ "wx_server_go/initial"

	_ "wx_server_go/routers"

	"fmt"
	"wx_server_go/models"

	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {

}

func main() {
	fmt.Println(time.Now().UnixNano())
	item, err := models.GetStatistics()
	fmt.Println(fmt.Sprintf("%+v", item), err)
	fmt.Println(time.Now().UnixNano())

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.BConfig.WebConfig.StaticDir["/static"] = "static"
	beego.SetStaticPath("/", "static")
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{beego.AppConfig.String("accessdomain")},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "token", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	//	beego.ErrorController(&web.ErrorController{})
	beego.Run()
}
