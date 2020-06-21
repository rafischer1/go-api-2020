package main

import (
	"github.com/gin-gonic/gin"
	"workspace/portfolio-go-api-2020/data"
)

func main() {
	r := gin.Default()
	r.GET("", func(c *gin.Context) {
		c.JSON(200, data.WelcomeMessage)
	})
	r.GET("/plans", func(c *gin.Context) {
		c.JSON(200, data.Plans)
	})
	r.GET("/companies", func(c *gin.Context) {
		c.JSON(200, data.Companies)
	})
	r.GET("/personal", func(c *gin.Context) {
		c.JSON(200, data.Personal)
	})
	_ = r.Run() // listen and serve on :8080
}
