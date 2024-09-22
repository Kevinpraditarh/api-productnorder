package repository

import (
	"api-productnorder/models"
	"database/sql"
	"time"
)

func GetAllProducts(db *sql.DB) ([]models.Datum, error) {
	rows, err := db.Query("SELECT id, name, price, sold, stock, created_at, updated_at FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Datum
	for rows.Next() {
		var product models.Datum
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Sold, &product.Stock, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func CreateProduct(db *sql.DB, name string, price int64, stock int64) (models.Data, error) {
	createdAt := time.Now().Format("2006-01-02 15:04:05")
	updatedAt := createdAt
	sold := int64(0)

	result, err := db.Exec("INSERT INTO products (name, price, stock, sold, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)", name, price, stock, sold, createdAt, updatedAt)
	if err != nil {
		return models.Data{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return models.Data{}, err
	}

	product := models.Data{
		ID:        id,
		Name:      name,
		Price:     price,
		Stock:     stock,
		Sold:      sold,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	return product, nil
}

func GetProductByID(db *sql.DB, id int64) (models.Data, error) {
	var product models.Data
	err := db.QueryRow("SELECT id, name, price, sold, stock, created_at, updated_at FROM products WHERE id = ?", id).Scan(
		&product.ID, &product.Name, &product.Price, &product.Sold, &product.Stock, &product.CreatedAt, &product.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Data{}, nil
		}
		return models.Data{}, err
	}
	return product, nil
}

func UpdateProduct(db *sql.DB, id int64, name string, price int64, stock int64) (models.Data, error) {
	updatedAt := time.Now().Format("2006-01-02 15:04:05")
	_, err := db.Exec("UPDATE products SET name = ?, price = ?, stock = ?, updated_at = ? WHERE id = ?", name, price, stock, updatedAt, id)
	if err != nil {
		return models.Data{}, err
	}

	return GetProductByID(db, id)
}

func DeleteProduct(db *sql.DB, id int64) (models.Data, error) {
	product, err := GetProductByID(db, id)
	if err != nil {
		return models.Data{}, err
	}

	if product.ID == 0 {
		return models.Data{}, nil
	}

	_, err = db.Exec("DELETE FROM products WHERE id = ?", id)
	if err != nil {
		return models.Data{}, err
	}

	return product, nil
}
