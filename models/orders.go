package models

type OrderDetailResponse struct {
	Message string      `json:"message"`
	Data    OrderDetail `json:"data"`
}

type CreateOrderResponse struct {
	Message string `json:"message"`
	Data    Order  `json:"data"`
}

// OrderDetail memuat detail informasi order dan produk yang terkait
type OrderDetail struct {
	ID        *int64    `json:"id"`
	Products  []Product `json:"products"`
	CreatedAt *string   `json:"created_at"`
	UpdatedAt *string   `json:"updated_at"`
}

type DetailOrder struct {
	Data    Order  `json:"data"`
	Message string `json:"message"`
}

type ListOrder struct {
	Data    []Order `json:"data"`
	Message string  `json:"message"`
}

type Order struct {
	CreatedAt *string   `json:"created_at,omitempty"`
	ID        *int64    `json:"id,omitempty"`
	Products  []Product `json:"products,omitempty"`
	UpdatedAt *string   `json:"updated_at,omitempty"`
}

type Product struct {
	CreatedAt string `json:"created_at"`
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Price     int64  `json:"price"`
	Quantity  int64  `json:"quantity"`
	Sold      int64  `json:"sold"`
	Stock     int64  `json:"stock"`
	UpdatedAt string `json:"updated_at"`
}

type DeleteOrder struct {
	Data    Order  `json:"data"`
	Message string `json:"message"`
}
