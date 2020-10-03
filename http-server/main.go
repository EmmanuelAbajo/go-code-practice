package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "URL.Path = %q\n", req.URL.Path)
		fmt.Fprintf(res, "%s %s %s\n", req.Method, req.URL, req.Proto)
		for k, v := range req.Header {
			fmt.Fprintf(res, "Header[%q] = %q\n", k, v)
		}
		fmt.Fprintf(res, "Host = %q\n", req.Host)
		fmt.Fprintf(res, "RemoteAddr = %q\n", req.RemoteAddr)
	})

	http.ListenAndServe("localhost:8000", nil)
}
