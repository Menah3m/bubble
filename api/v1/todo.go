package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 代办事项

// 查看
func GetTodo(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{
		"status": "GetTodo",
	})
}
// 增加
func AddTodo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "AddTodo",
	})
}


// 修改
func EditTodo(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{
		"status": "EditTodo",
	})
}
// 删除
func DeleteTodo(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{
		"status": "DeleteTodo",
	})
}
