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
	defer db.Close()
}
