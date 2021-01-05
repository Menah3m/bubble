package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

)

var (
	DB *gorm.DB
)

// 连接数据库

func InitMySQL() (err error)  {
	dsn := "root:root1234@tcp(192.168.13.110:13306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)

	return
}