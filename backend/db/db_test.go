package db

import (
	"testing"
)

func TestInitDB(t *testing.T) {
	db := InitDB()
	if db == nil {
		t.Errorf("InitDB() returned nil")
	}
}
