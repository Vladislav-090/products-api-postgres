package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"product-api-postgres/internal/models"
	"product-api-postgres/internal/response"
	"product-api-postgres/internal/storage"
	"strconv"
)

type ProductHandler struct {
	Storage *storage.ProductStorage
}

func NewProductHandler(productStorage *storage.ProductStorage) *ProductHandler {
	return &ProductHandler{
		Storage: productStorage,
	}
}

func (h *ProductHandler) AddProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.WriteError(w, http.StatusMethodNotAllowed, "Method not allowed!")
		return
	}

	var product models.Product

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, "Invalid JSON!")
		return
	}

	if product.Title == "" {
		response.WriteError(w, http.StatusBadRequest, "Title is empty!")
		return
	}

	if product.Price <= 0 {
		response.WriteError(w, http.StatusBadRequest, "Price must be positive!")
		return
	}

	createProduct, err := h.Storage.CreateProduct(product)
	if err != nil {
		log.Println("failed create product!", err)
		response.WriteError(w, http.StatusInternalServerError, " Failed to create product!")
		return
	}

	response.WriteSucces(w, http.StatusCreated, "Product Created Succsessfully!", createProduct)
}

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response.WriteError(w, http.StatusMethodNotAllowed, "Method not allowed!")
		return
	}

	products, err := h.Storage.GetProducts()
	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, "Failed to get products")
		return
	}

	response.WriteJSON(w, http.StatusOK, products)
}

func (h ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response.WriteError(w, http.StatusMethodNotAllowed, "Method not allowed!")
		return
	}

	idParam := r.URL.Query().Get("id")
	if idParam == "" {
		response.WriteError(w, http.StatusBadRequest, "ID is empty")
		return
	}
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, "Error to convert string to int!")
		return
	}

	product, err := h.Storage.GetProduct(id)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, "product not found!")
		return
	}

	response.WriteJSON(w, http.StatusOK, product)
}
