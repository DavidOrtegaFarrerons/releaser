package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func NewPostgresClient() (*sql.DB, error) {
	//Hardcoded values just for testing purposes, will be moved to .env
	host := "localhost"
	port := "5432"
	user := "appuser"
	password := "apppassword"
	dbname := "appdb"

	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %w", err)
	}

	// Optional: test the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("db.Ping: %w", err)
	}

	return db, nil
}
