package db

import (
	"testing"
)

func TestOpenDB(t *testing.T) {

	client, err := NewPostgresClient()
	if err != nil {
		t.Fatalf("failed to create DB client: %v", err)
	}
	defer client.Close()

	if err := client.Ping(); err != nil {
		t.Fatalf("failed to ping DB: %v", err)
	}

	t.Log("✅ Connected successfully using NewPostgresClient()")
}
