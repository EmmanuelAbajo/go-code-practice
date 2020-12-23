package main

import (
	"fmt"
	"time"
)

//BankAccount struct
type BankAccount struct {
	firstName string
	lastName string
	bvn string
	balance float32
	typeOfAccount string
	accountNumber string
}

// Employee struct
type Employee struct {
	id int
	firstName string
	lastName string
	createdAt time.Time
}

func (emp Employee) setName(firstName, lastName string) {
	emp.firstName = firstName
	emp.lastName = lastName
}

func (emp Employee) getIDByFirstName(firstName string) int {
	return emp.id
}

func (emp Employee) toString() {
	fmt.Printf("id=%d firstname=%s lastname=%s createdAt=%s\n", emp.id, emp.firstName, emp.lastName, emp.createdAt)
}


func main () {
	var newEmp Employee
	emp := new(Employee)

	emp.id = 1
	emp.firstName = "John"
	emp.lastName = "Doe"
	emp.createdAt = time.Now()

	newEmp.id = 2
	newEmp.firstName = "Jack"
	newEmp.lastName = "Dorsey"
	newEmp.createdAt = time.Now()

	

	fmt.Println(emp.createdAt)
	newEmp.toString()
	fmt.Println(newEmp.getIDByFirstName("jack"))
	
}