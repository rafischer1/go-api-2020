package handlers

import (
	"fmt"
)

// Email (params: address string from Url) start email service
func Email(address string) {
	fmt.Println("in the email service!:", address)
}
