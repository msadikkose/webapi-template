package controllers

import (
	"webapi-template/managers"

	"github.com/kataras/iris"
)

//GetDateTimeNow is returns Datetime.Now as JSON
func GetDateTimeNow(c *iris.Context) {
	c.JSON(iris.StatusOK, managers.GetDateTimeNow())
}

//SayHello is returns greating
func SayHello(c *iris.Context) {
	name := c.Param("name")
	c.JSON(iris.StatusOK, managers.SayHello(name))
}
