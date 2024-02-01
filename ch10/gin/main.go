package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// go modules 添加依赖
// 删除未使用的依赖项
// go get, go mod ? go install
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
