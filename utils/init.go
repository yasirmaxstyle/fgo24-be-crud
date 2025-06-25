package utils

import (
	"context"
	"log"
)

func InitDB() {
	conn := DBConnect()
	defer func() {
		conn.Conn().Close(context.Background())
	}()

	// Create table if not exists
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS contacts (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL,
		phone VARCHAR(20) NOT NULL,
		created_at TIMESTAMP DEFAULT now(),
		updated_at TIMESTAMP DEFAULT now()
	);
	
	-- Create index for better search performance
	CREATE INDEX IF NOT EXISTS idx_contacts_name ON contacts(name);
	`

	if _, err := conn.Exec(context.Background(), createTableQuery); err != nil {
		log.Fatal("Failed to create table:", err)
	}

	log.Println("Database connected and table created successfully")
}
