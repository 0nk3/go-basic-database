package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// User : how do we display the data from the db ?
type User struct {
	id        int    `json:"id"`
	firstname string `json:"firstname"`
	lastname  string `json:"lastname"`
}

func main() {
	log.Println("initialise the driver and potentially open the DB")
	database, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/user_database")
	if err != nil {
		panic(err.Error())
	}
	defer database.Close()
	log.Println("database successfully created and connected to")
	// display
	result, err := database.Query("SELECT * from users")
	if err != nil {
		panic(err.Error())
	}
	for result.Next() {
		var user User
		err = result.Scan(&user.id, &user.firstname, &user.lastname)
		if err != nil {
			panic(err.Error())
		}
		log.Println(user.id, user.firstname, user.lastname)
	}

	//Lets create the table
	// insert, err := database.Query("INSERT INTO users(id, firstname,lastname) VALUES (4, 'Dr Kabutz', 'Heinz')")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer insert.Close()
	// log.Println("User successfully insertted into the the db of users table")
	// delete, err := database.Query("DELETE FROM users WHERE firstname='tim'")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// log.Println("time remove")
	// defer delete.Close()
}
