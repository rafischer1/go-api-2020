package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"workspace/portfolio-go-api-2020/data"
	"workspace/portfolio-go-api-2020/handlers"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("templates/index.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": data.WelcomeMessage,
		})
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
