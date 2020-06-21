package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"workspace/portfolio-go-api-2020/handlers"
)

func main() {
	r := gin.Default()
	r.GET("", func(c *gin.Context) {
		handlers.GetAll("home", c)
	})
	r.GET("/plans", func(c *gin.Context) {
		handlers.GetAll("plans", c)
	})
	r.GET("/companies", func(c *gin.Context) {
		handlers.GetAll("companies", c)
	})
	r.GET("/contact", func(c *gin.Context) {
		handlers.GetAll("contact", c)
	})
	err := r.Run() // listen and serve on :8080
	if err != nil {
		fmt.Println("Error on listen and serve:", err)
	}
}
