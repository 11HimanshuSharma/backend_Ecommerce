package database


import (
	"database/sql"
	"fmt"
	"log"
	"time"

	// the underscore "_" means we import it for side effects (registering the driver)
	// but we dont call its functions directly in this file
	_ "github.com/lib/pq"
)



/// Config holds our database connection details
type Config struct {
	Host string
	Port string
	User string
	Password string
	DBName string
	SSLMode string
}

// connect initializes a postgresSQL connection pool and returns the DB instance
func Connect(cfg Config) (*sql.DB, error) {
	//// build the data source name (DSN) connection sring
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode,
	)

	// openthe database connecton just validating arg
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)

	}

	// configure connection pool limits
	// dont legt go overwhelm postres
	db.SetMaxOpenConn(25) // max simultaneous connection to DB
	db.SetMaxIdleConn(25) // keep these open when idle to save overhead
	db.SetConnMaxLifetime(5 * time.Minute) // Force refreshing connection after 5 m


	// ping to verify the connection is actually alive
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}
	log.Println("Successfully connected to postgreSQL database")
	return db, nil
}


