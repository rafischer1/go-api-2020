package main

import (
	"fmt"
	"net/http"
	"os"
	d "workspace/portfolio-go-api-2020/data"
	h "workspace/portfolio-go-api-2020/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("env in main.go: ", getEnv())
	r := gin.Default()

	r.LoadHTMLFiles("templates/index.html", "templates/resume.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": d.WelcomeMessage,
		})
	})
	r.GET("/resume", func(c *gin.Context) {
		c.HTML(http.StatusOK, "resume.html", gin.H{})
	})
	r.GET("/companies", func(c *gin.Context) {
		h.GetAll("companies", c, "")
	})
	r.GET("/contact", func(c *gin.Context) {
		h.GetAll("contact", c, "")
	})
	r.GET("/plans", func(c *gin.Context) {
		h.GetAll("plans", c, "")
	})
	r.GET("/email", func(c *gin.Context) {
		h.GetAll("email", c, c.Query("mailTo"))
	})
	r.GET("/Robert_Fischer_Resume_2020", func(c *gin.Context) {
		h.GetAll("download", c, "")
	})

	r.Use(cors.Default())
	err := r.Run() // listen and serve on :8080
	if err != nil {
		fmt.Print("Error on r.run():", err)
	}
}

func getEnv() string {
	return os.Getenv("APP_ENV")
}
