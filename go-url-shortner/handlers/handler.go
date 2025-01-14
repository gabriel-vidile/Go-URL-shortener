package handlers

import (
	"encoding/json"
	"go-url-shortner/storage"
	"go-url-shortner/utils"
	"net/http"
	"sync"
)

type Request struct {
	URL string `json:"url"`
}

type Response struct {
	ShortURL string `json:"short_url"`
}

var redirectCache sync.Map

func ShortenURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Requisição inválida", http.StatusBadRequest)
		return
	}

	if existingShortID, found := storage.GetShortID(req.URL); found {
		resp := Response{ShortURL: "http://localhost:8080/" + existingShortID}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
		return
	}

	shortID := utils.GenerateShortID()
	err := storage.Save(shortID, req.URL)
	if err != nil {
		http.Error(w, "Erro ao salvar URL", http.StatusInternalServerError)
		return
	}

	resp := Response{ShortURL: "http://localhost:8080/" + shortID}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
func RedirectURL(w http.ResponseWriter, r *http.Request) {
	shortID := r.URL.Path[1:]

	if cachedURL, found := redirectCache.Load(shortID); found {
		http.Redirect(w, r, cachedURL.(string), http.StatusFound)
		return
	}

	originalURL, err := storage.Get(shortID)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	redirectCache.Store(shortID, originalURL)

	http.Redirect(w, r, originalURL, http.StatusFound)
}
