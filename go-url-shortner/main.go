package main

import (
	"go-url-shortner/handlers"
	"go-url-shortner/storage"
	"log"
	"net/http"
)

func main() {
	err := storage.InitializeDatabase("go_url_shortener.db")
	if err != nil {
		log.Fatalf("Erro ao inicializar banco de dados: %v", err)
	}
	defer storage.CloseDatabase()

	http.HandleFunc("/shorten", handlers.ShortenURL)
	http.HandleFunc("/", handlers.RedirectURL)

	log.Println("Servidor iniciado na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
