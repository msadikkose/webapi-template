package routers

import (
	"github.com/kataras/iris"
	"github.com/msadikkose/webapi-template/controllers"
)

//SetHomeRoute sets the home routes
func SetHomeRoute(api *iris.Application) *iris.Application {

	api.Get("/home/", controllers.GetDateTimeNow)
	api.Get("/home/sayhello/{name:string min(1)}", controllers.SayHello)

	return api
}
