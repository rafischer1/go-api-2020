package handlers

import (
	"github.com/gin-gonic/gin"
	"workspace/portfolio-go-api-2020/data"
)

// GetAll switch statement for all static data handlers params:(path string, context)
func GetAll(path string, c *gin.Context, params string) {
	c.Header("Access-Control-Allow-Origin", "*")
	switch path {
	case "plans":
		c.JSON(200, data.Plans)
	case "companies":
		c.JSON(200, data.Companies)
	case "contact":
		c.JSON(200, data.Personal)
	case "download":
		Download(params)
	case "email":
		Email(params)
	}
}
