package main

import (
	"github.com/gin-gonic/gin"
	"workspace/portfolio-go-api-2020/data"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/plans", func(c *gin.Context) {
		c.JSON(200, data.Plans)
	})
	r.GET("/companies", func(c *gin.Context) {
		c.JSON(200, data.Companies)
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
