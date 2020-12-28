package main

import (
	"github.com/Menah3m/bubble/routers"
	"github.com/gin-gonic/gin"
)


func main()  {
	gin.DisableConsoleColor()
	r := routers.InitRouter()
	r.Run()
}
