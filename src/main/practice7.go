package main

import (
	"fmt"
	_"github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
)

func main() {
	fmt.Println("START")

	cnn, _ := sql.Open("mysql", "root:test@tcp(127.0.0.1:3306)/TestDB")

	var title string
	if err := cnn.QueryRow("SELECT Title FROM Category WHERE id = ? LIMIT 1", 1).Scan(&title); err != nil {
		log.Fatal(err)
	}
	fmt.Println(title)

	fmt.Println("END")
}
