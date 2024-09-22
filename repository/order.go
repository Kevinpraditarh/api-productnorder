package repository

import (
	"api-productnorder/models"
	"database/sql"
)

// GetOrders retrieves a list of orders with their related products.
func GetOrders(db *sql.DB) ([]models.Order, error) {
	// Query to get orders
	rows, err := db.Query(`
		SELECT o.id, o.created_at, o.updated_at, p.id, p.name, p.price, op.quantity, p.stock, p.sold, p.created_at, p.updated_at
		FROM orders o
		JOIN order_products op ON o.id = op.order_id
		JOIN products p ON op.product_id = p.id
		ORDER BY o.id;
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ordersMap := make(map[int64]*models.Order)

	for rows.Next() {
		var orderID int64
		var orderCreatedAt, orderUpdatedAt string
		var product models.Product

		err := rows.Scan(&orderID, &orderCreatedAt, &orderUpdatedAt,
			&product.ID, &product.Name, &product.Price, &product.Quantity,
			&product.Stock, &product.Sold, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			return nil, err
		}

		// Check if order already exists in the map
		if _, ok := ordersMap[orderID]; !ok {
			ordersMap[orderID] = &models.Order{
				ID:        &orderID,
				CreatedAt: &orderCreatedAt,
				UpdatedAt: &orderUpdatedAt,
				Products:  []models.Product{},
			}
		}

		// Append product to the order
		ordersMap[orderID].Products = append(ordersMap[orderID].Products, product)
	}

	// Convert map to slice
	var orders []models.Order
	for _, order := range ordersMap {
		orders = append(orders, *order)
	}

	return orders, nil
}

func CreateOrder(db *sql.DB, products []models.Product) (int64, error) {
	// Simpan order ke dalam tabel orders
	result, err := db.Exec("INSERT INTO orders (created_at, updated_at) VALUES (NOW(), NOW())")
	if err != nil {
		return 0, err
	}

	orderID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	// Simpan produk terkait order di tabel order_products
	for _, product := range products {
		_, err := db.Exec("INSERT INTO order_products (order_id, product_id, quantity, created_at, updated_at) VALUES (?, ?, ?, NOW(), NOW())",
			orderID, product.ID, product.Quantity)
		if err != nil {
			return 0, err
		}
	}

	return orderID, nil
}

func GetOrderByID(db *sql.DB, id int64) (models.Order, error) {
	var order models.Order

	query := `SELECT id, created_at, updated_at FROM orders WHERE id = ?`
	err := db.QueryRow(query, id).Scan(&order.ID, &order.CreatedAt, &order.UpdatedAt)
	if err != nil {
		return models.Order{}, err
	}

	// Dapatkan produk terkait
	order.Products, err = GetProductsByOrderID(db, id)
	if err != nil {
		return models.Order{}, err
	}

	return order, nil
}

// Fungsi untuk mendapatkan produk berdasarkan order ID
func GetProductsByOrderID(db *sql.DB, orderID int64) ([]models.Product, error) {
	var products []models.Product
	query := `SELECT p.id, p.name, p.price, op.quantity, p.stock, p.sold, p.created_at, p.updated_at
			  FROM order_products op
			  JOIN products p ON op.product_id = p.id
			  WHERE op.order_id = ?`

	rows, err := db.Query(query, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Quantity, &product.Stock, &product.Sold, &product.CreatedAt, &product.UpdatedAt); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func DeleteOrderByID(db *sql.DB, id int64) error {
	query := `DELETE FROM orders WHERE id = ?`
	_, err := db.Exec(query, id)
	return err
}
