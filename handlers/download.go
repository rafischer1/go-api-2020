package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"io"
	"os"
)

// Download redirects toPDF served url - takes gin.Context and contains error handling
func Download(c *gin.Context) {
	// Open file
	_ = godotenv.Load(".env")
	f, err := os.Open(os.Getenv("RESUME_DEV_URL"))
	if err != nil {
		fmt.Println(err)
		ErrorMsg(500, err, c)
		return
	}
	defer f.Close()

	//Set header
	c.Header("Content-type", "application/pdf")
	//Stream to response
	if _, err := io.Copy(c.Writer, f); err != nil {
		fmt.Println(err)
		ErrorMsg(500, err, c)
	}
}

func ErrorMsg(code int, err error, c *gin.Context) {
	c.String(code, "Error formatting PDF load")
	// A custom error page with HTML templates can be shown by c.HTML()
	_ = c.Error(err)
	c.Abort()
}
