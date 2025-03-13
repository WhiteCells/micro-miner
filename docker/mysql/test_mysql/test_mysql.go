package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "testuser:testpassword@tcp(127.0.0.1:3309)/testdb"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("sql Open error:", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("sql ping error:", err)
	}

	fmt.Println("success open MySQL")
}
