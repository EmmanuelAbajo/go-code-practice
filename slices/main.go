package main

import (
	"fmt"
	"time"
)

// 1. Slices are dynamic while arrays are not
//
// 2. As slices are passed by reference to functions, which means that what is actually passed is
// the memory address of the slice variable, any modifications you make to a slice inside a
// function will not get lost after the function exits.
//
// 3.Additionally, passing a big slice to a function is significantly faster than passing
// an array with the same number of elements because Go will not have to make a copy of the slice;
// it will just pass the memory address of the slice variable
//
// 4.  This means that altering the elements of a reslice modifies the element of the
// original slice because they both point to the same underlying array
func main() {
	start := time.Now()

	arr1 := [5]int{1, 2, 3, 4, 5}  // Go array
	arr2 := [...]string{1: "January", 2: "February", 3: "March"} // Array with indices
	slice1 := make([]int, 10)      // Go slice
	slice2 := []int{1, 2, 3, 4, 5} // Go slice
	slice3 := arr1[1:3]            // Slice from array
	slice4 := slice1[0:5]          // Go reslicing

	fmt.Println(arr1)
	fmt.Println(arr2[2])
	fmt.Println(slice3)
	slice2 = append(slice2, 6)
	fmt.Println(slice2)
	slice1[0] = 10
	slice1 = append(slice1, 20)
	fmt.Println(slice1)

	fmt.Println(slice4)
	k := []int{100,200,300}
	slice4 = append(slice4,k...)
	fmt.Println(slice4)

	duration := time.Since(start)
	fmt.Println("Execution time:", duration)

}
