package controllers

import (
	"github.com/kataras/iris"
	"github.com/msadikkose/webapi-template/managers"
)

//GetDateTimeNow is returns Datetime.Now as JSON
func GetDateTimeNow(c iris.Context) {
	c.JSON(managers.GetDateTimeNow())
	}

//SayHello is returns greating
func SayHello(c iris.Context) {

	name := c.Params().Get("name")
	c.JSON(managers.SayHello(name))

}

