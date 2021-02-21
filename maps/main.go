package main

import "fmt"
import "sort"

func testPrintFnc(value interface{}) {
	fmt.Println(value)
}

func main() {

	var x map[string]int
	x = make(map[string]int)
	if val, ok := x["John"]; !ok {
		fmt.Println("Key not found")
	} else {
		fmt.Println(val)
	}

	x["John"] = 20
	x["Joe"] = 15
	x["Brad"] = 25

	names := make([]string, 0, len(x))
	
	for name := range x {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, x[name])
	}

	delete(x, "John")
	testPrintFnc(x)


}
