package main

import (
	"log"
	"fmt"
	"os"
	"strings"
	"./tempconv"
	"github.com/mactsouk/go/simpleGitHub"
)

var greeting string // Variable declaration
var cwd string

func main()  {
	var err error
	cwd,err = os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
		os.Exit(1)
	}
	log.Printf("Working directory = %s\n", cwd)

	greeting = "Hello world" // Assignment
	var freezingTemp, boilingTemp float64 = 32.0, 212.0

	fmt.Println(*newInt());
	fmt.Println(greeting)
	fmt.Println(tempconv.AbsoluteZeroC)
	fmt.Printf("%goF = %goT\n",freezingTemp,fToC(freezingTemp))
	fmt.Printf("%goF = %goT\n",boilingTemp,fToC(boilingTemp))

	count := 5 // short variable declaration and initialization
	for i := 0; i < count; i++ {
		fmt.Println(i)
	}

	for index, value := range [5]int{1,2,3,4,5} {
		fmt.Printf("index: %d, value: %d\n",index,value)
	}

	fmt.Println(strings.Join(os.Args[1:]," "))
	fmt.Println(simpleGitHub.AddTwo(5, 6))
}

func fToC(f float64) float64  {
	var t = ( f- 32 ) * 5/9
	return t
}

func newInt() *int {
	var n int
	return &n
}