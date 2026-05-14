package repository

import (
	"database/sql"
	"fmt"
	"ecommerce/internal/models"
)


type postgresOrderRepo struct {
	db *sql.DB

}

func NewPostgresOrderRepo(db *sql.DB) OrderRepository {
	return &postgresOrderRepo{
		db: db,
	}

}

func (r *postgresOrderRepo) CreateOrder(order *models.Order) error {
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	// Perform order creation logic here
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}