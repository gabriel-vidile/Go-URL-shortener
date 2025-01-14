package tests

import (
	"bytes"
	"encoding/json"
	"go-url-shortner/handlers"
	"go-url-shortner/storage"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func setupHandlersDatabase(t *testing.T) {
	err := storage.InitializeDatabase("test.db")
	if err != nil {
		t.Fatalf("Erro ao inicializar banco de dados: %v", err)
	}
}

func teardownHandlersDatabase() {
	storage.CloseDatabase()
	os.Remove("test.db")
}

func TestShortenURLHandler(t *testing.T) {
	setupHandlersDatabase(t)
	defer teardownHandlersDatabase()

	requestBody, _ := json.Marshal(map[string]string{
		"url": "https://example.com",
	})
	req := httptest.NewRequest(http.MethodPost, "/shorten", bytes.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.ShortenURL)

	handler.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("Esperava status 200, recebeu %d", recorder.Code)
	}

	var response map[string]string
	json.Unmarshal(recorder.Body.Bytes(), &response)
	shortURL, ok := response["short_url"]
	if !ok || shortURL == "" {
		t.Errorf("Resposta inválida: %s", recorder.Body.String())
	}
}
