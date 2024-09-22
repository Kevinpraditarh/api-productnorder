package models

type ApidogModel struct {
	Data    []Datum `json:"data"`
	Message string  `json:"message"`
}

type Datum struct {
	CreatedAt string `json:"created_at"`
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Price     int64  `json:"price"`
	Sold      int64  `json:"sold"`
	Stock     int64  `json:"stock"`
	UpdatedAt string `json:"updated_at"`
}

type CreateProduct struct {
	Data    Data   `json:"data"`
	Message string `json:"message"`
}

type DetailProduct struct {
	Data    Data   `json:"data"`
	Message string `json:"message"`
}

type Data struct {
	CreatedAt string `json:"created_at"`
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Price     int64  `json:"price"`
	Sold      int64  `json:"sold"`
	Stock     int64  `json:"stock"`
	UpdatedAt string `json:"updated_at"`
}
