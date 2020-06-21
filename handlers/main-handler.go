package handlers

import (
	"github.com/gin-gonic/gin"
	"workspace/portfolio-go-api-2020/data"
)

// GetAll switch statement for all static data handlers params:(path string, context)
func GetAll(path string, c *gin.Context) {
	switch path {
	case "home":
		c.JSON(200, data.WelcomeMessage)
	case "plans":
		c.JSON(200, data.Plans)
	case "companies":
		c.JSON(200, data.Companies)
	case "contact":
		c.JSON(200, data.Personal)
	}
}