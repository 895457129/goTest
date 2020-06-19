package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"ranyun/goTest/controllers"
)


func main() {
	app := iris.New()
	app.Done(func(ctx iris.Context) {})
	mvcApi := mvc.New(app).Party("/api")
	mvcApi.Party("/user").Handle(new(controllers.OrderController))
	_ = app.Run(iris.Addr(":8080"))
}


