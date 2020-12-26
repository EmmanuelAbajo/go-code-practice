package main

import (
	"fmt"
	"time"
	"os"
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

func initBankAccount(firstName string,
	lastName string,
	bvn string,
	balance float32,
	typeOfAccount string,
	accountNumber string)  (*BankAccount, error) {

		if len(accountNumber) < 10 {
			return nil, fmt.Errorf("Account number must be >= 10 not %s", accountNumber)
		}

		if balance < 0 {
			return nil, fmt.Errorf("Initial balance must be >= 0 not %d", balance)
		}

		account := &BankAccount { 
			firstName, 
			lastName, 
			bvn, 
			balance, 
			typeOfAccount, 
			accountNumber, 
		}
		
		return account, nil
}

// we use a pointer receiver so this method won't get a copy of the struct
// but a reference to the original struct
func (acc *BankAccount) creditAccount(amount float32) {
	acc.balance += amount
}

func (acc *BankAccount) withdraw(amount float32) {
	acc.balance -= amount
}

func (acc *BankAccount) checkBalance(amount float32) float32{
	return acc.balance
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

	testAcc, err := initBankAccount("John", "Michael", "", 2000, "savings", "1000035786")
	if err != nil {
		fmt.Printf("Error in creating account - %s\n", err)
		os.Exit(1)
	}

	fmt.Println(*testAcc)
	
}