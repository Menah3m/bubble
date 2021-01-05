package v1

import (
	"github.com/Menah3m/bubble/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 所有代办事项
func GetAllTodo(c *gin.Context)  {
	// 查询Todo表里的所有数据
	var todoList []models.Todo
	if err := models.DB.Find(&todoList).Error; err != nil{
		c.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todoList)
	}

}
// 查看某项待办事项
func GetTodo(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{
		"status": "GetTodo",
	})
}
// 增加
func AddTodo(c *gin.Context) {
	// 前端页面填写请求 点击提交 会发请求到这里
	// 从请求中拿出数据
	var todo models.Todo
	c.BindJSON(&todo)

	//存入数据库
	if err := models.DB.Create(&todo).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	}else {
		c.JSON(http.StatusOK, todo)
	}
	// 返回 结果
	c.JSON(http.StatusOK, gin.H{
		"status": "AddTodo",
	})
}


// 修改
func EditTodo(c *gin.Context)  {
	id, ok := c.Params.Get("id")
	if !ok{
		c.JSON(http.StatusOK, gin.H{
			"error": "无效的id",
		})
		return
	}
	var todo models.Todo
	if err := models.DB.Where("id=?",id).First(&todo).Error; err != nil{
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	}
	c.BindJSON(&todo)
	if err := models.DB.Save(&todo).Error; err != nil{
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todo)
	}


}
// 删除
func DeleteTodo(c *gin.Context)  {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"error": "无效的id",
		})
		return
	}
	var todo models.Todo
	if err := models.DB.Where("id=?", id).Delete(&todo).Error;err != nil{
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "已删除",
		})
	}

}
