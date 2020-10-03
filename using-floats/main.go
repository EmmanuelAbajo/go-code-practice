package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	var n float64
	var err error = errors.New("New Error Found")
	
	arguments := os.Args
	k := 1
	if len(arguments) == 1 {
		io.WriteString(os.Stdout, "Provide one or more floats"+"\n")
		os.Exit(1)
	}
	
	for err != nil {
		if k >= len(arguments) {
			fmt.Println("None of the arguments is a float!")
			return
		}
		n, err = strconv.ParseFloat(arguments[k], 64)
		if err != nil {
			fmt.Println(err.Error())
		}
		k++
	}
	min, max := n, n

	for i := 2; i < len(arguments); i++ {
		n, err := strconv.ParseFloat(arguments[i], 64)
		if err == nil {
			if n < min {
				min = n
			}
			if n > max {
				max = n
			}
		}
	}

	fmt.Println("Min:", min)
	fmt.Println("Max:", max)

}
