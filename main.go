package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	d "workspace/portfolio-go-api-2020/data"
	h "workspace/portfolio-go-api-2020/handlers"
)

func main() {
	r := gin.Default()

	r.LoadHTMLFiles("templates/index.html", "templates/resume.html")

	r.GET("/", func(c *gin.Context) {
		setHeader(c)
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": d.WelcomeMessage,
		})
	})
	r.GET("/resume", func(c *gin.Context) {
		setHeader(c)
		c.HTML(http.StatusOK, "resume.html", gin.H{})
	})
	r.GET("/companies", func(c *gin.Context) {
		setHeader(c)
		h.GetAll("companies", c)
	})
	r.GET("/contact", func(c *gin.Context) {
		setHeader(c)
		h.GetAll("contact", c)
	})
	r.GET("/plans", func(c *gin.Context) {
		setHeader(c)
		h.GetAll("plans", c)
	})

	r.Use(cors.Default())
	err := r.Run() // listen and serve on :8080
	if err != nil {
		fmt.Print("Error on r.run():", err)
	}
}
func setHeader(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
}
