package main

import (
	"api-productnorder/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/products", handlers.GetProductsHandler).Methods("GET")
	r.HandleFunc("/api/products", handlers.CreateProductHandler).Methods("POST")
	r.HandleFunc("/api/products/{id:[0-9]+}", handlers.GetProductDetailHandler).Methods("GET")
	r.HandleFunc("/api/products/{id:[0-9]+}", handlers.UpdateProductHandler).Methods("PUT")
	r.HandleFunc("/api/products/{id:[0-9]+}", handlers.DeleteProductHandler).Methods("DELETE")

	r.HandleFunc("/api/orders", handlers.GetOrdersHandler).Methods("GET")
	r.HandleFunc("/api/orders", handlers.CreateOrderHandler).Methods("POST")
	r.HandleFunc("/api/orders/{id}", handlers.GetOrderDetailHandler).Methods("GET")
	r.HandleFunc("/api/orders/{id}", handlers.DeleteOrderHandler).Methods("DELETE")

	http.Handle("/", r)
	fmt.Println("Terhubung ke server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
