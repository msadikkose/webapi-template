package main

import (

	"github.com/kataras/iris"
	"github.com/msadikkose/webapi-template/common"
	"github.com/msadikkose/webapi-template/routers"
)

func main() {

	common.InitConfig()
	app := iris.New()
	routers.SetHomeRoute(app)


	app.Run(iris.Addr(common.AppConfig.Port))
}
