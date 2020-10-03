package main

import (
	"os"
	"bufio"
	"io"
	"fmt"
)

func main()  {
	var f *os.File;
	f = os.Stdin;
	defer f.Close();

	fmt.Println("Hi there, enter text. Enter ctrl + D to stop reading from the std input");
	scanner := bufio.NewScanner(f);
	for scanner.Scan() {
		io.WriteString(os.Stdout,">>>" + scanner.Text() + "\n");
	}
}