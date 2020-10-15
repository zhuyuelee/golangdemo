package main

import (
	"fmt"
	"net/http"
)

var urls = []string{
	"http://www.baidu.com/",
	"http://www.baidu.com/bac/",
	"https://www.cnblogs.com/yanghongtao/",
	"https://www.cnblogs.com/yanghongtao/p/11160253.html",
}

func main() {
	// Execute an HTTP HEAD request for all url's
	// and returns the HTTP status string or an error string.
	for _, url := range urls {
		resp, err := http.Head(url)
		if err != nil {
			fmt.Println("Error:", url, err)
		}
		fmt.Println(url, ": ", resp.Status)
	}
}
