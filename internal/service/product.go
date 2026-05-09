package service

import "ecommerce/internal/models"


// ProductService defines business logic for operations on products.
// service is the "brain" of our application, it applies business rules.
type ProductService interface {
	// getProduct retrieves a product by its ID, return an error
	GetProduct(id int64) (*models.Product, error)

	// ListProducts returns all available products
	ListProducts() ([]*models.Product, error)


	//AddProduct creates a new product and it validates that price is positive
	AddProduct(name, description string, price float64, stock int32) (*models.Product, error) 

	//DeductStock reduces the stock of a product when an order is placed
	// this is a core business rule, we muust check stock before deducint
	DeductStock(productID int64, quantity int32) error 
}