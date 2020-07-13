package handlers

import (
	"fmt"
	"github.com/cavaliercoder/grab"
	"log"
)

// Email (params: address string from Url) start email service
func Email(address string) {
	fmt.Println("in the email service!:", address)
	url := "https://github.com/rafischer1/go-api-2020/blob/master/assets/Robert_Arthur_Fischer_Test_Resume_2020.pdf"
	resp, err := grab.Get(".", url)
	if err != nil {
		log.Fatal(err)
	}

	//http.Post("download", "Content-type:application/pdf", resp.Filename)
	fmt.Println("Download saved to", resp.Filename)
}
