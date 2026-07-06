package main

import (
	"fmt"
	"log"
	"product-api-postgres/internal/database"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Println("Connected to database succsessfully!")
}