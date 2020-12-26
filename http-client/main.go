package main

import (
	"net/http"
	"os"
	"fmt"
	"io/ioutil"
)

func main(){
	url := os.Args[1]
	response,err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	
	defer response.Body.Close()
	fmt.Println("Response code:",response.Status)

	// returns a byte array
	data,err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	fmt.Printf("%s\n", data)

}