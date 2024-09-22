package handlers

import (
	"api-productnorder/config"
	"api-productnorder/models"
	"api-productnorder/repository"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func GetOrdersHandler(w http.ResponseWriter, r *http.Request) {
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

	orders, err := repository.GetOrders(db)
	if err != nil {
		http.Error(w, "Failed to retrieve orders", http.StatusInternalServerError)
		return
	}

	response := models.ListOrder{
		Data:    orders,
		Message: "Order List",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
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
		Products []struct {
			ID       int64 `json:"id"`
			Quantity int64 `json:"quantity"`
		} `json:"products"`
	}

	err = json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var orderProducts []models.Product
	for _, productReq := range requestBody.Products {
		product, err := repository.GetProductByID(db, productReq.ID)
		if err != nil {
			log.Println("Product not found:", productReq.ID) // Log error
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}

		if product.Stock < productReq.Quantity {
			log.Println("Product out of stock for product ID:", productReq.ID) // Log error
			http.Error(w, "Product out of stock", http.StatusBadRequest)
			return
		}

		product.Stock -= productReq.Quantity
		product.Sold += productReq.Quantity

		_, err = repository.UpdateProduct(db, product.ID, product.Name, product.Price, product.Stock)
		if err != nil {
			fmt.Println("Failed to update product stock for product ID:", product.ID) // Log error
			http.Error(w, "Failed to update product stock", http.StatusInternalServerError)
			return
		}

		orderProduct := models.Product{
			ID:        product.ID,
			Name:      product.Name,
			Price:     product.Price,
			Quantity:  productReq.Quantity,
			Sold:      product.Sold,
			Stock:     product.Stock,
			CreatedAt: product.CreatedAt,
			UpdatedAt: product.UpdatedAt,
		}

		orderProducts = append(orderProducts, orderProduct)
	}

	orderID, err := repository.CreateOrder(db, orderProducts)
	if err != nil {
		fmt.Println("Failed to create order:", err) // Log error
		http.Error(w, "Failed to create order", http.StatusInternalServerError)
		return
	}

	currentTime := time.Now().Format(time.RFC3339)
	response := models.DetailOrder{
		Data: models.Order{
			ID:        &orderID,
			Products:  orderProducts,
			CreatedAt: &currentTime,
			UpdatedAt: &currentTime,
		},
		Message: "Order created",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func GetOrderDetailHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/api/orders/"):]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	db, err := config.ConnectDB()
	if err != nil {
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	order, err := repository.GetOrderByID(db, id)
	if err != nil {
		http.Error(w, "Failed to retrieve order", http.StatusInternalServerError)
		return
	}

	if order.ID == nil {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	response := models.DetailOrder{
		Data:    order,
		Message: "Order Detail",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func DeleteOrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/api/orders/"):]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	db, err := config.ConnectDB()
	if err != nil {
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	order, err := repository.GetOrderByID(db, id)
	if err != nil || (order.ID != nil && *order.ID == 0) {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	err = repository.DeleteOrderByID(db, id)
	if err != nil {
		http.Error(w, "Failed to delete order", http.StatusInternalServerError)
		return
	}

	response := models.DeleteOrder{
		Data:    order,
		Message: "Order deleted successfully",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
