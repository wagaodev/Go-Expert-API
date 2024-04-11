package database

import (
	"testing"

	"github.com/wagaodev/Go-Expert-API/internal/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupTestDB(t *testing.T) *gorm.DB {
	t.Helper() // Indica que esta é uma função de ajuda

	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Falha ao abrir o banco de dados: %v", err)
	}

	if err := db.AutoMigrate(&entity.Product{}); err != nil {
		t.Fatalf("Falha na migração do banco de dados: %v", err)
	}

	return db
}
