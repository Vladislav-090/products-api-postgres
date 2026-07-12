package main

import (
	"fmt"
	"log"
	"net/http"
	"product-api-postgres/internal/database"
	"product-api-postgres/internal/handlers"
	"product-api-postgres/internal/storage"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Println("Connected to database succsessfully!")

	productStorage := storage.NewProductStorage(db)
	productHandler := handlers.NewProductHandler(productStorage)

	http.HandleFunc("/addProduct", productHandler.AddProductHandler)
	http.HandleFunc("/getProducts", productHandler.GetProductsHandler)
	http.HandleFunc("/getProduct", productHandler.GetProductHandler)
	http.HandleFunc("/updateProduct", productHandler.UpdateProductHandler)
	http.HandleFunc("/deleteProduct", productHandler.DeleteProductHandler)
	http.HandleFunc("/getProductsCount", productHandler.GetCountHandler)

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
