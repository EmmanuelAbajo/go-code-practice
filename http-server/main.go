package main

import (
	"net/http"
	"./controller"
)

func main() {
	http.HandleFunc("/", controller.GetURLDetails)
	http.HandleFunc("/gif", controller.GetLissajous)
	http.ListenAndServe("localhost:8000", nil)
}




