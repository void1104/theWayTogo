package main

import (
	"fmt"
	"net/http"
)

var urls = []string{
	"https://www.baidu.com",
	"https://github.com",
	"https://www.alibaba.com",
}

func main() {
	// Execute an HTTP HEAD request for all urls
	// and returns the HTTP status string or an error string.
	for _, url := range urls {
		resp, err := http.Head(url)
		if err != nil {
			fmt.Println("Error:", url, err)
		}
		fmt.Println(url, ": ", resp.Status)
	}
}
