package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// main 简单的模拟gin服务。
func main() {
	engine := gin.New()
	engine.Handle("GET", "/", func(context *gin.Context) {
		context.String(200, "success")
		fmt.Println("nothing")
	})
	engine.Run(":9999")
}
