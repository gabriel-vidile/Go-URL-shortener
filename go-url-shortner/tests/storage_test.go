package tests

import (
	"go-url-shortner/storage"
	"os"
	"testing"
)

func setupDatabase(t *testing.T) {
	err := storage.InitializeDatabase("test.db")
	if err != nil {
		t.Fatalf("Erro ao inicializar banco de dados: %v", err)
	}
}

func teardownDatabase() {
	storage.CloseDatabase()
	os.Remove("test.db")
}

func TestSaveAndGet(t *testing.T) {
	setupDatabase(t)
	defer teardownDatabase()

	originalURL := "https://example.com"
	shortID := "abcd1234"

	err := storage.Save(shortID, originalURL)
	if err != nil {
		t.Errorf("Erro ao salvar URL: %v", err)
	}

	retrievedURL, err := storage.Get(shortID)
	if err != nil || retrievedURL != originalURL {
		t.Errorf("Erro ao recuperar URL: esperava %s, recebeu %s", originalURL, retrievedURL)
	}
}

func TestGetShortID(t *testing.T) {
	setupDatabase(t)
	defer teardownDatabase()

	originalURL := "https://example.com"
	shortID := "abcd1234"

	err := storage.Save(shortID, originalURL)
	if err != nil {
		t.Errorf("Erro ao salvar URL: %v", err)
	}

	retrievedShortID, found := storage.GetShortID(originalURL)
	if !found || retrievedShortID != shortID {
		t.Errorf("Erro ao recuperar shortID: esperava %s, recebeu %s", shortID, retrievedShortID)
	}
}
