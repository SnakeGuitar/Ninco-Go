package db

import (
	"os"
	"testing"
	"time"
)

func TestCreateConnection(t *testing.T) {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	if host == "" || user == "" {
		t.Skip("Skipping integration test: DB_HOST and DB_USER environment variables not set")
	}

	cfg := Config{
		User:            user,
		Password:        pass,
		Host:            host,
		Port:            port,
		Database:        name,
		MaxIdleConns:    5,
		MaxOpenConns:    10,
		ConnMaxLifetime: time.Hour,
	}

	db, err := CreateConnection(cfg)
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	defer db.Close()

	if err := Ping(db); err != nil {
		t.Errorf("Database is reachable but Ping failed: %v", err)
	}
}
