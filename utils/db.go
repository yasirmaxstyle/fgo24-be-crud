package utils

import (
	"context"
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

	pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Failed to create connection pool:", err)
	}

	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.Fatal("Failed to create connection from the pool:", err)
	}
	return conn
}
