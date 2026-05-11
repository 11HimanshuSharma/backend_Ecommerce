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

	query := `
	   SELECT id, name, description, price, stock, created_at, updated_at
	   FROM products
	   WHERE id = $1
	`

	p := &models.Product{}

	err := r.db.QueryRow(query, id).Scan(
		&p.ID,
		&p.Name,
		&p.Description,
		&p.Price,
		&p.Stock,
		&p.CreatedAt,
		&p.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("product with id %d not found", id)
		}
		return nil, fmt.Errorf("error fetching product: %w", err)
	}
	return p, nil
}

func (r *postgresProductRepo) GetAll() ([]*models.Product, error) {
	query := `
	SELECT id, name, description, price, stock, created_at, updated_at
	FROM products
	ORDER BY created_at DESC	
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error fetching products: %w", err)
	}
	defer rows.Close() // always defer rows.Close() immediately after sql.DB.Query() to prevent resource leaks

	var products []*models.Product
	for rows.Next() {
		p := &models.Product{}
		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Description,
			&p.Price,
			&p.Stock,
			&p.CreatedAt,
			&p.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning product: %w", err)
		}
		products = append(products, p)
	}

	return products, nil
}


func (r *postgresProductRepo) Create(product *models.Product) (int64, error) {
	query := `
	INSERT INTO products (name, description, price, stock, created_at, updated_at)
	VALUES ($1, $2, $3, $4, NOW(), NOW())
	RETURNING id
	   `

	var newID int64
	// In PostgreSQL, we use RETURNING to grab the newly generated ID directly
	err := r.db.QueryRow(
		query, 
		product.Name,
		product.Description,
		product.Price,
		product.Stock,
	).Scan(&newID)

	if err != nil {
		return 0, fmt.Errorf("error creating product: %w", err)

	}
	return newID, nil
}


func (r *postgresProductRepo) Update(product *models.Product) error {
	query := `
	UPDATE products
	SET name = $1, description = $2, price = $3, stock = $4, updated_at = NOW()
	WHERE id = $5
	`
	result, err := r.db.Exec(
		query, 
		product.Name,
		Product.Description,
		product.Price,
		product.Stock,
		product.ID,
	)
	if err != nil {
		return fmt.Errorf("error updating product: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking update result: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("product with id %d not found", product.ID)
	}

	return nil
}

func (r *postgresProductRepo) Delete(id int64) error {
	query := `
	DELETE FROM products WHERE id = $1
	`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting product: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking delete result: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("product with id %d not found", id)
	}

	return nil
	}