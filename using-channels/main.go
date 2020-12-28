// Get content type of sites
package main

import (
	"fmt"
	"net/http"
)

// receive string from http call
// publish result to string channel
// return channel 
// in main, print out result from channel
func returnContentType(url string, ch chan string) {
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("error: %s\n", err)
		return
	}

	defer resp.Body.Close()
	ctype := resp.Header.Get("content-type")
	ch <- fmt.Sprintf("%s -> %s", url, ctype)
	
}

func main() {
	ch := make(chan string)
	urls := []string{
		"https://golang.org",
		"https://api.github.com",
		"https://httpbin.org/xml",
	}

	for _, url := range urls {
		// run go routine
		go returnContentType(url, ch)
	}

	for range urls {
		res := <- ch
		fmt.Println(res)
	}
}
