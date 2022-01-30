package handlers

import (
	"fmt"
	"workspace/portfolio-go-api-2020/data"

	"github.com/gin-gonic/gin"
)

// GetAll switch statement for all static data handlers params:(path string, context, params?)
func GetAll(path string, c *gin.Context, params string) {
	c.Header("Access-Control-Allow-Origin", "*")
	switch path {
	case "plans":
		c.JSON(200, data.Plans)
	case "companies":
		c.JSON(200, data.Companies)
	case "trek":
		c.JSON(200, data.TrekBooks)
	case "contact":
		c.JSON(200, data.Personal)
	case "download":
		Download(c)
	case "email":
		Email(c, params)
	}
}

// GetByID switch statement for all static data handlers params:(path string, context, params?)
func GetByID(path string, c *gin.Context, params string) {
	c.Header("Access-Control-Allow-Origin", "*")
	switch path {
	case "trek/:id":
		res := []data.TrekBook{}
		for i := range data.TrekBooks {
			if fmt.Sprint(data.TrekBooks[i].ID) == params {
				res = append(res, data.TrekBooks[i])
			}
		}
		c.JSON(200, res)
	}
}
