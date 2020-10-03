package main

import (
	"bufio"
	"fmt"
	"os"
)

func countLines(f *os.File, count map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		count[input.Text()]++
	}
}

func main() {
	fileName := "animals.txt"
	counts := make(map[string]int)
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	countLines(f, counts)
	f.Close()

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
