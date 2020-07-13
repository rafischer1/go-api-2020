package handlers

import (
	"fmt"
	"github.com/cavaliercoder/grab"
	"log"
)

// Email (params: address string from Url) start email service
func Email(address string) {
	fmt.Println("in the email service!:", address)
	resp, err := grab.Get(".", "http://www.golang-book.com/public/pdf/gobook.pdf")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Download saved to", resp.Filename)
}
