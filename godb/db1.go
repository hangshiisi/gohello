package main

import (
	"fmt"
	//"math/rand"
	//"time"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql",
		"root:cisco123@tcp(127.0.0.1:3306)/godb")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Hello World Connection Success \n")

	err = db.Ping()
	if err != nil {
		// do something here
		fmt.Print("Error after Connection Success \n")

	}

	var (
		name  string
		owner string
	)

	//rows, err := db.Query("select  name, owner from pet where id = ?", 1)
	rows, err := db.Query("select  name, owner from pet")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&name, &owner)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(name, owner)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}
