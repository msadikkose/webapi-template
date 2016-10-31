package main

import (
	"webapi-template/common"
	"webapi-template/routers"

	"github.com/kataras/iris"
)

func main() {
	common.InitConfig()
	api := iris.New()
	routers.SetHomeRoute(api)
	api.Listen(common.AppConfig.Port)
}
