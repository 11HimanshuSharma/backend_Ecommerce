package service

import (
	"ecommerce/internal/models"
	"ecommerce/internal/repository"
	"ecommerce/pkg/apperror"
)

// prpductServiceImpl holds out business logic and talks to any repository
type productServiceImpl struct {
	// Notic we use the interface here, not the struct
	// this is Dependency injection
	repo repository.ProductRepository
}

// NewProductservice injects the repository into our service layer
func NewProductService(r repository.ProductRepository) ProductService {
	return &productServiceImpl{
		repo: r,
	}
}

func (s *productServiceImpl) GetProduct(id int64) (*models.Product, error) {
	return s.repo.GetByID(id)
}
func (s *productServiceImpl) ListProducts() ([]*models.Product, error) {
	return s.repo.GetAll()
}

func (s *productServiceImpl) AddProduct(name, description string, price float64, stock int) (*models.Product, error) {
	// valid price
	if price <= 0 {
		return nil, apperror.NewBadRequestError("INVALID_PRICE", "Price must be positive")
	}
	// valid stock
	if stock < 0 {
		return nil, apperror.NewBadRequestError("INVALID_STOCK", "Stock cannot be negative")
	}
	product := &models.Product{
		Name:        name,
		Description: description,
		Price:       price,
		Stock:       stock,
	}
	// tell the repository to save tis
	id, err := s.repo.Create(product)
	if err != nil {
		return nil, err
	}

	// fetch the newly created product to return to the user
	return s.repo.GetByID(id)
}

func (s *productServiceImpl) DeductStock(productId int64, quantity int) error {
	if quantity <= 0 {
		return apperror.NewBadRequestError("INVALID_QUANTITY", "Quantity must be positive")
	}
	product, err := s.repo.GetByID(productId)
	if err != nil {
		return err
	}
	if product.Stock < quantity {
		return apperror.NewConflict("INSUFFICIENT_STOCK", "Not enough stock available")
	}
	product.Stock -= quantity
	return s.repo.Update(product)
}
