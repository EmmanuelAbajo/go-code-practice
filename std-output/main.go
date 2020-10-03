package main

import (
	"os"
	"io"
)

func main(){
	arguments := os.Args;
	if len(arguments) == 1 {
		io.WriteString(os.Stdout,"Provide one argument"+ "\n");
	} else {
		io.WriteString(os.Stdout, arguments[1] + "\n");
	}

}