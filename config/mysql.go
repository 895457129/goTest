package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

var MysqlDb *gorm.DB

func init() {
	fmt.Println("初始化数据库")
	var err error
	MysqlDb, err = gorm.Open("mysql", "develop:develop123@/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	// 最大连接数
	MysqlDb.DB().SetMaxOpenConns(100)
	// 闲置连接数
	MysqlDb.DB().SetMaxIdleConns(20)
	// 最大连接周期
	MysqlDb.DB().SetConnMaxLifetime(100 * time.Second)
}



