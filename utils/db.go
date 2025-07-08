package utils

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func DBConnect() *pgxpool.Conn {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file")
	}
	connectionString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGDATABASE"),
	)

	pool, err := pgxpool.New(context.Background(), connectionString)
	if err != nil {
		log.Fatal("Failed to create connection pool:", err)
	}

	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.Fatal("Failed to create connection from the pool:", err)
	}
	return conn
}
