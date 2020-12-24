// print numbers 1 - 20
// if number is divisible by 3, print fizz
// if number is divisible by 5, print buzz
// If number is divisible by both, print fizzbuzz
package main

import "fmt"

func main() {
	for i := 1; i <= 20; i++ {
		switch {
		case (i%3 == 0 && i%5 == 0):
			fmt.Println("fizzbuzz")
			break
		case (i%3 == 0):
			fmt.Println("fizz")
			break
		case (i%5 == 0):
			fmt.Println("buzz")
			break
		default:
			fmt.Println(i)
			break
		}
	}
}
