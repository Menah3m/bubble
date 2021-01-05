package main

import (
	"github.com/Menah3m/bubble/routers"
	"github.com/Menah3m/bubble/models"
	"github.com/gin-gonic/gin"
)


func main()  {
	gin.DisableConsoleColor() // 关闭console颜色显示

	err := models.InitMySQL() // 连接数据库
	if err != nil{
		panic(err)
	}
	defer models.DB.Close() // 程序退出时关闭数据库

	models.DB.AutoMigrate(&models.Todo{})  // 建立数据模型

	r := routers.InitRouter()  // 初始化路由
	r.Run()
}
