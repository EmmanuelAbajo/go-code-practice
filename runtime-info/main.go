package main

import (
	"fmt"
	"runtime"
)

// defer uses LIFO i.e last function to be deferred is the first function
// to be implemented. defer pushes the call of a function to just before the
// end of its surrounding function returns
func main() {
	details := make([]*int, 10) 
	
	fmt.Println("=====Starting Function Execution=====")
	defer fmt.Println("=====Ending Function Execution=====")

	defer func() {
		fmt.Println("Architecture target:", runtime.GOARCH)
		fmt.Println("Go version:", runtime.Version())
		fmt.Println("Number of CPUs:", runtime.NumCPU())
	}()

	for i := range details {
		details[i] = &i
		// This is an anonymous function
		func(value *int) {
			fmt.Println(value)
			fmt.Printf("Type of value: %T\n",value)
		}(details[i])

	}

}
