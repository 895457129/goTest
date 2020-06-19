package controllers

import (
	"github.com/kataras/iris"
	"ranyun/goTest/config"
	"ranyun/goTest/models"
)

type OrderController struct {}

func (u *OrderController) Post(ctx iris.Context) string {
	var order models.Order
	_ = ctx.ReadJSON(&order)
	if !config.MysqlDb.HasTable(&models.Order{}) {
		config.MysqlDb.CreateTable(&models.Order{})
	}
	config.MysqlDb.Create(&order)
	return "操作成功"
}