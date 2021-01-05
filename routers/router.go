package routers

import (
	v1 "github.com/Menah3m/bubble/api/v1"
	"net/http"

	"github.com/gin-gonic/gin"

)

func InitRouter() *gin.Engine  {
	r := gin.Default()
	// 告诉gin框架模板文件引用的静态文件在哪里
	r.Static("/static", "static")
	// 导入模板文件
	r.LoadHTMLFiles("./templates/index.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	apiV1 := r.Group("v1")
	{
		apiV1.GET("/todo/", v1.GetAllTodo)
		apiV1.GET("/todo/:id", v1.GetTodo)
		apiV1.POST("/todo", v1.AddTodo)
		apiV1.PUT("/todo/:id", v1.EditTodo)
		apiV1.DELETE("/todo/:id", v1.DeleteTodo)
	}

	return r
}
