package repository

import (
	"errors"
	"sync"
	"time"

	"ecommerce/internal/models"
)

// InMemoryProductRepository implements productRepository
type InMemoryProductRepo struct {
	// A map to store products, and a mutex to prevent data races
	// when multiple HTTP requests try to read// write at the same time
	mu sync.RWMutex
	products map[int64]*models.Product
	nextID int64
}



//NewInMemoryProductRepo is a constructor function
func NewInMemoryProductRepo() *InMemoryProductrepo {
	return &InMemoryProductRepo{
		products: make(map[int64]*models.Product),
		nextID: 1,
	}

}

//getById fethces a product by itd is
func (r *InMemoryProductRepo) GetByID(id int64) (*models.Product, error) {
	// read lock
	r.mu.RLock()
	defer r.mu.RUnlock()

	product, exists := r.products[id]
	if !exists {
		return nil, errors.New("Product not found")
	}
	return product, nil
}


// getAll return all products
func (r *InMemoryProductRepo) GetAll() ([]*models.Product, error){
	r.mu.RLock()
	defer r.mu.RUnlock()

	var all []*models.Product
	for _, p := range r.products {
		all = append(all, p)
	}
	return all, nil
}


//create generates a new ID, sets timestamps, and saves the product
func (r *InMemoryProductRepo) Create(product *models.Product) (int64, error) {
	r.mu.Lock()// write lock
	defer r.mu.Unlock()

	product.ID = r.nextID
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	r.products[product.ID] = product
	r.nextID++

	return product.ID, nil
}
	
// Update modifies an existing product in the map
func (r *InMemoryProductRepo) Update(product *models.Product) error {
	r.mu.Lock()

	defer r.mu.Unlock()

	if _, exists := r.products[product.ID]; !exists {
		return errors.New("Product not found")
	}
	product.UpdatedAt = time.Now()
	r.products[product.ID] = product
	return nil
}

func (r *InMemoryProductRepo) Delete(id int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.products, id)
	return nil
}