package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

// PostTrekBook handles the submit form
func PostTrekBook(path string, c *gin.Context, params string) {
	c.Header("Access-Control-Allow-Origin", "*")
	// Read the Body content
	var formDataBytes []byte
	if c.Request.Body != nil {
		formDataBytes, _ = ioutil.ReadAll(c.Request.Body)
	}

	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(string(formDataBytes)), &jsonMap)

	fmt.Println(jsonMap)
	c.Next()
	c.JSON(200, "Post Successfull")
}
