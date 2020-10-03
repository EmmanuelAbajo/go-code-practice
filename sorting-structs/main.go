package main

import (
	"fmt"
	"sort"
)

// Person represents a person object
type Person struct {
	name string
	age int
}

func main() {
	personSlice := make([]Person,0)
	personSlice = append(personSlice,Person{"Rob",23})
	personSlice = append(personSlice,Person{"Max",24})
	personSlice = append(personSlice,Person{"Jack",25})
	personSlice = append(personSlice,Person{"Bob",21})

	fmt.Println(personSlice)
	
	sort.Slice(personSlice,func(i,j int) bool {
		return personSlice[i].age > personSlice[j].age
	})
	fmt.Println("Decending:",personSlice)

	sort.Slice(personSlice,func(i,j int) bool {
		return personSlice[i].age < personSlice[j].age
	})
	fmt.Println("Ascending:",personSlice)

	// Using new returns a pointer
	personSlice2 := new(Person)
	personSlice2.name = "Ryan"
	personSlice2.age = 20
	fmt.Println(*personSlice2)

}