package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
	"strings"
)

func main() {
	id := ulid.MustNew(0, strings.NewReader("ABCDEFGHIJKLMNOP"))
	fmt.Println("------UUID:>",id.String())
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}