package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// PostTrekBook handles the submit form
func PostTrekBook(path string, c *gin.Context, params string) {
	c.Header("Access-Control-Allow-Origin", "*")
	fmt.Println(params)
	c.JSON(200, "Post Successfull")
}
