package controller

import (
	"database/sql"
	"fmt"

	"github.com/Utsavch189/api_mysql/internal/config"
	"github.com/Utsavch189/api_mysql/internal/models/requests"
)

func AddProduct(product *requests.Product) (*requests.Product, error) {
	query := `INSERT INTO Product (id, product_name, price, created_at, is_active) 
	          VALUES (?, ?, ?, ?, ?)`

	db, err := config.Connect()

	if err != nil {
		return &requests.Product{}, err
	}
	defer db.Close()
	_, qerr := db.Exec(query, product.Id, product.ProductName, product.Price, product.CreatedAt, product.IsActive)

	if qerr != nil {
		return &requests.Product{}, qerr
	}
	return product, nil
}

func GetProducts() ([]requests.Product, error) {
	var products []requests.Product
	db, err := config.Connect()

	if err != nil {
		return products, err
	}
	defer db.Close()

	query := `Select id, product_name, price, created_at, is_active From Product`
	rows, err := db.Query(query)
	if err != nil {
		return products, err
	}
	defer rows.Close()

	for rows.Next() {
		var product requests.Product
		err := rows.Scan(
			&product.Id,
			&product.ProductName,
			&product.Price,
			&product.CreatedAt,
			&product.IsActive,
		)
		if err != nil {
			return products, err
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return products, err
	}

	return products, nil
}

func GetProduct(id string) (requests.Product, error) {
	var product requests.Product

	db, err := config.Connect()

	if err != nil {
		return product, err
	}
	defer db.Close()

	query := `Select id, product_name, price, created_at, is_active From Product where id=?`
	row := db.QueryRow(query, id)
	err = row.Scan(
		&product.Id,
		&product.ProductName,
		&product.Price,
		&product.CreatedAt,
		&product.IsActive,
	)
	if err == sql.ErrNoRows {
		return product, fmt.Errorf("product with id %s not found", id)
	}
	if err != nil {
		return product, err
	}
	return product, nil
}

func UpdateProduct(product *requests.Product) (*requests.Product, error) {
	db, err := config.Connect()

	if err != nil {
		return &requests.Product{}, err
	}
	defer db.Close()

	query := `Update Product Set product_name=?, price=? Where id=?`
	result, cerr := db.Exec(query, product.ProductName, product.Price, product.Id)

	if cerr != nil {
		return &requests.Product{}, cerr
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return &requests.Product{}, fmt.Errorf("error fetching rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return &requests.Product{}, fmt.Errorf("no product found with id: %s", product.Id)
	}

	return product, nil
}
func DeleteProduct(id string) error {
	db, err := config.Connect()

	if err != nil {
		return err
	}
	defer db.Close()

	query := `Delete from Product Where id=?`
	result, cerr := db.Exec(query, id)

	if cerr != nil {
		return cerr
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error fetching rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no product found with id: %s", id)
	}

	return nil
}
