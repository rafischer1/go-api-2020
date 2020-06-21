package main

import (
	"github.com/gin-gonic/gin"
	"workspace/portfolio-go-api-2020/data"
)

func main() {
	r := gin.Default()
	r.GET("", func(c *gin.Context) {
		homePath(c)
	})
	r.GET("/plans", func(c *gin.Context) {
		plans(c)
	})
	r.GET("/companies", func(c *gin.Context) {
		companies(c)
	})
	r.GET("/contact", func(c *gin.Context) {
		contact(c)
	})
	_ = r.Run() // listen and serve on :8080
}

func homePath(c *gin.Context) {
	c.JSON(200, data.WelcomeMessage)
}

func plans(c *gin.Context) {
	c.JSON(200, data.Companies)
}

func companies(c *gin.Context) {
	c.JSON(200, data.Plans)
}

func contact(c *gin.Context) {
	c.JSON(200, data.Personal)
}
