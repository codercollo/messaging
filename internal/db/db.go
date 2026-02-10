package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Database struct {
	Client *sqlx.DB
}

func NewDatabase() (*Database, error) {
	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_TABLE"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("SSL_MODE"),
	)

	var dbConn *sqlx.DB
	var err error

	// Retry connection for up to 30 seconds
	for i := 0; i < 15; i++ {
		dbConn, err = sqlx.Connect("postgres", connectionString)
		if err == nil {
			fmt.Println("Successfully connected to database")
			break
		}

		fmt.Printf("Database not ready, retrying in 2 seconds... (attempt %d/15)\n", i+1)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		return &Database{}, fmt.Errorf("could not connect to the database after retries: %w", err)
	}

	return &Database{
		Client: dbConn,
	}, nil
}

func (d *Database) Ping(ctx context.Context) error {
	return d.Client.DB.PingContext(ctx)
}
