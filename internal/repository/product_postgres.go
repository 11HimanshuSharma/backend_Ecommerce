package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"ecommerce/internal/models"
)

type postgresProductRepo struct {
	db *sql.DB
}

func NewPostgresProductRepo(db *sql.DB) ProductRepository {
	return &postgresProductRepo{db: db}
}

func (r *postgresProductRepo) GetByID(id int64) (*models.Product, error) {
	
}