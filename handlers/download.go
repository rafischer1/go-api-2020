package handlers

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Reader interface {
	Read(buf []byte) (n int, err error)
}

//// Download (params: timestamp from url) start download service
//func Download(address string) {
//	fmt.Println("in the email service!:", address)
//	url := "https://github.com/rafischer1/go-api-2020/blob/master/assets/Robert_Arthur_Fischer_Test_Resume_2020.pdf"
//	//resp, err := grab.Get(".", url)
//	//if err != nil {
//	//	log.Fatal(err)
//	//}
//
//	//http.Post("download", "Content-type:application/pdf", resp.Filename)
//	r := strings.NewReader(url)
//	_, _ = http.Post("http://localhost:8080/download", "application/pdf", r)
//	file, _ := os.Create("Test_Resume.pdf")
//	os.OpenFile(url)
//	//fmt.Println("Download saved to", resp.Filename, response)
//}
var (
	fileName    string
	fullUrlFile string
)

func Download() {
	fullUrlFile = "https://github.com/rafischer1/go-api-2020/blob/master/assets/Robert_Arthur_Fischer_Test_Resume_2020.pdf"

	// Build fileName from fullPath
	buildFileName()

	// Create blank file
	file := createFile()

	// Put content on file
	putFile(file, httpClient())

}

func putFile(file *os.File, client *http.Client) {
	resp, err := client.Get(fullUrlFile)

	checkError(err)

	defer resp.Body.Close()

	size, err := io.Copy(file, resp.Body)

	defer file.Close()

	checkError(err)

	fmt.Printf("Just Downloaded a file %s with size %d\n", fileName, size)
}

func buildFileName() {
	fileUrl, err := url.Parse(fullUrlFile)
	checkError(err)

	path := fileUrl.Path
	segments := strings.Split(path, "/")

	fileName = segments[len(segments)-1]
}

func httpClient() *http.Client {
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	return &client
}

func createFile() *os.File {
	file, err := os.Create(fileName)

	checkError(err)
	return file
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
