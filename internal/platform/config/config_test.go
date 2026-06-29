package config

import (
	"os"
	"testing"
)

func TestLoadDB_Defaults(t *testing.T) {
	for _, k := range []string{
		"DB_HOST", "DB_PORT", "DB_USER", "DB_PASSWORD", "DB_NAME", "DB_SSLMODE",
	} {
		os.Unsetenv(k)
	}

	db := LoadDB()

	if db.Host != "localhost" || db.Port != "54321" || db.User != "pangea" {
		t.Fatalf("Unexpected defaults: %v", db)
	}

	want := "postgres://pangea:pangea@localhost:54321/pangea?sslmode=disable"

	if got := db.URL(); got != want {
		t.Errorf("URL Mismatch: \n got: %q\nwant: %q", got, want)
	}
}

func TestLoadDB_Override(t *testing.T) {
	os.Setenv("DB_HOST", "db.internal")
	defer os.Unsetenv("DB_HOST")

	if got := LoadDB().Host; got != "db.internal" {
		t.Errorf("Expected override, got: %q", got)
	}
}
