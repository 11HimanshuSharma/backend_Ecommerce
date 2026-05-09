package reposi

import "ecommerce/internal/models"



// product repository defines the contract for product data access.
// any implementation must procide these methods
type ProductRepository interface {

	GetByID(id int64) (*models.Product, error)

	// getAll returns all products 
	GetAll() ([]*models.Product, error)


	//create insert a new product and return its ID.
	Create(product *models.Product) (int64, error)

	// Update modifies an existing product
	Update(product *models.Product) error

	// Delete remove a product by its ID
	Delete(id int64) error
	
}
