package handlers

import (
	"api-productnorder/config"
	"api-productnorder/models"
	"api-productnorder/repository"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	db, err := config.ConnectDB()
	if err != nil {
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	products, err := repository.GetAllProducts(db)
	if err != nil {
		http.Error(w, "Failed to retrieve products", http.StatusInternalServerError)
		return
	}

	if len(products) == 0 {
		http.Error(w, "No products found", http.StatusNotFound)
		return
	}

	response := models.ApidogModel{
		Data:    products,
		Message: "Products retrieved successfully",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// CreateProductHandler handles POST requests to create a new product
func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	db, err := config.ConnectDB()
	if err != nil {
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var requestBody struct {
		Name  string `json:"name"`
		Price int64  `json:"price"`
		Stock int64  `json:"stock"`
	}

	err = json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	product, err := repository.CreateProduct(db, requestBody.Name, requestBody.Price, requestBody.Stock)
	if err != nil {
		fmt.Printf("Error creating product: %v", err)
		http.Error(w, "Failed to create product", http.StatusInternalServerError)
		return
	}

	response := models.CreateProduct{
		Data:    product,
		Message: "Product created successfully",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// GetProductDetailHandler handles GET requests for a single product by ID
func GetProductDetailHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/api/products/"):] // Get ID from URL
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	db, err := config.ConnectDB()
	if err != nil {
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	product, err := repository.GetProductByID(db, id)
	if err != nil {
		http.Error(w, "Failed to retrieve product", http.StatusInternalServerError)
		return
	}

	if product.ID == 0 {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	response := models.DetailProduct{
		Data:    product,
		Message: "Product Detail",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// UpdateProductHandler handles PUT requests to update a product
func UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/api/products/"):] // Get ID from URL
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	db, err := config.ConnectDB()
	if err != nil {
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var requestBody struct {
		Name  string `json:"name"`
		Price int64  `json:"price"`
		Stock int64  `json:"stock"`
	}

	err = json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	product, err := repository.UpdateProduct(db, id, requestBody.Name, requestBody.Price, requestBody.Stock)
	if err != nil {
		http.Error(w, "Failed to update product", http.StatusInternalServerError)
		return
	}

	if product.ID == 0 {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	response := models.DetailProduct{
		Data:    product,
		Message: "Product updated successfully",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// DeleteProductHandler handles DELETE requests to delete a product
func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/api/products/"):] // Get ID from URL
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	db, err := config.ConnectDB()
	if err != nil {
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	product, err := repository.DeleteProduct(db, id)
	if err != nil {
		http.Error(w, "Failed to delete product", http.StatusInternalServerError)
		return
	}

	if product.ID == 0 {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	response := struct {
		Message string      `json:"message"`
		Data    models.Data `json:"data"`
	}{
		Message: "Product deleted successfully",
		Data:    product,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
