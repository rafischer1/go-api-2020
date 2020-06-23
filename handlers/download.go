package handlers

import (
	"fmt"
)

// Download (params: timestamp from url) start download service
func Download(timestamp string) {
	fmt.Println("in the download function!:", timestamp)
}
