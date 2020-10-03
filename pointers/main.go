package main

import "fmt"

func getPointer(n *int) {
	*n = *n * *n
}

func returnPointer(n int) *int {
	v := n * n
	return &v
}

func main() {
	i := 2
	j := 3
	pI := &i
	pJ := &j

	fmt.Println("pI memory:", pI)
	fmt.Println("pJ memory:", pJ)
	fmt.Println("pI value:", *pI)
	fmt.Println("pJ value:", *pJ)

	getPointer(pI)
	fmt.Println("i:", i)
	k := returnPointer(j)
	fmt.Println("k:", *k)
	fmt.Println("k memory: ", k)
}
