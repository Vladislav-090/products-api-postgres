package main

import (
	"fmt"
	"log"
	"net/http"
	"product-api-postgres/internal/database"
	"product-api-postgres/internal/handlers"
	"product-api-postgres/internal/middleware"
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
	userStorage := storage.NewUserStorage(db)
	authHandler := handlers.NewAuthHandler(userStorage)

	http.HandleFunc("/register", authHandler.Register)
	http.HandleFunc("/login", authHandler.Login)

	http.Handle(
		"/addProduct",
		middleware.AuthMiddleware(
			middleware.AdminMiddleware(
				http.HandlerFunc(productHandler.AddProductHandler)),
		),
	)

	http.Handle(
		"/getProducts",
		middleware.AuthMiddleware(
			http.HandlerFunc(productHandler.GetProductsHandler),
		),
	)
	http.Handle("/getProduct",
		middleware.AuthMiddleware(
			http.HandlerFunc(productHandler.GetProductHandler),
		),
	)
	http.Handle("/updateProduct",
		middleware.AuthMiddleware(
			middleware.AdminMiddleware(
				http.HandlerFunc(productHandler.UpdateProductHandler)),
		),
	)
	http.Handle("/deleteProduct",
		middleware.AuthMiddleware(
			middleware.AdminMiddleware(
				http.HandlerFunc(productHandler.DeleteProductHandler),
			),
		),
	)
	http.Handle("/getProductsCount",
		middleware.AuthMiddleware(
			http.HandlerFunc(productHandler.GetCountHandler),
		),
	)
	http.Handle("/clearProducts",
		middleware.AuthMiddleware(
			middleware.AdminMiddleware(
				http.HandlerFunc(productHandler.ClearProductsHandler),
			),
		),
	)

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
