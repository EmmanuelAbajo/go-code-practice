package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func insert(db *sql.DB) {
	insert, err := db.Query("INSERT INTO USERS VALUES(1,'Emmanuel','John')")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()

	fmt.Println("Successfully inserted to SQL db")
}

func read(db *sql.DB) {
	var user User

	result, err := db.Query("SELECT * FROM USERS")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err := result.Scan(&user.ID, &user.FirstName, &user.LastName)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(user)
	}
	defer result.Close()
}

func readOne(db *sql.DB, id int) {
	var user User

	err := db.QueryRow("SELECT * FROM USERS WHERE ID=?", id).Scan(&user.ID, &user.FirstName, &user.LastName)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(user)
}

func main() {
	fmt.Println("Connecting to SQL db...")

	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:6604)/testdb")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	fmt.Println("Successfully connected to SQL db")

	insert(db)
	read(db)
	readOne(db,1)

}
