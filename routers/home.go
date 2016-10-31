package routers

import (
	"webapi-template/controllers"

	"github.com/kataras/iris"
)

//SetHomeRoute sets the home routes
func SetHomeRoute(api *iris.Framework) *iris.Framework {

	api.Get("/home/", controllers.GetDateTimeNow)
	api.Post("/home/SayHello/:name", controllers.SayHello)

	return api
}
