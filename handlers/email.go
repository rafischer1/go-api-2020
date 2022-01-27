package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

// Email (params: address string from Url) start email service
func Email(c *gin.Context, address string) {
	_ = godotenv.Load(".env")
	f, err := os.Open(os.Getenv("RESUME_DEV_URL"))
	if err != nil {
		fmt.Println(err)
		ErrorMsg(500, err, c)
		return
	}
	fmt.Println("F:", f, "Address: ", address)
	defer f.Close()
}
