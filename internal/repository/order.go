package repository


import "ecommerce/internal/models"

type OrderRepository interface {
	CreateOrder(order *models.Order) error
}