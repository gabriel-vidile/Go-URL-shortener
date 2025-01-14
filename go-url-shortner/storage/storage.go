package storage

import (
	"database/sql"
	"errors"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

var (
	db    *sql.DB
	cache sync.Map
)

func InitializeDatabase(filePath string) error {
	var err error
	db, err = sql.Open("sqlite3", filePath)
	if err != nil {
		return err
	}

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS urls (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		short_id TEXT UNIQUE NOT NULL,
		original_url TEXT NOT NULL
	);`
	_, err = db.Exec(createTableQuery)
	return err
}

func CloseDatabase() {
	if db != nil {
		db.Close()
	}
}

func Save(shortID, originalURL string) error {
	_, err := db.Exec("INSERT INTO urls (short_id, original_url) VALUES (?, ?)", shortID, originalURL)
	if err == nil {
		cache.Store(originalURL, shortID)
	}
	return err
}

func GetShortID(originalURL string) (string, bool) {
	if cachedID, ok := cache.Load(originalURL); ok {
		return cachedID.(string), true
	}

	var shortID string
	err := db.QueryRow("SELECT short_id FROM urls WHERE original_url = ?", originalURL).Scan(&shortID)
	if err == sql.ErrNoRows {
		return "", false
	} else if err != nil {
		return "", false
	}

	cache.Store(originalURL, shortID)
	return shortID, true
}

func Get(shortID string) (string, error) {
	var originalURL string
	err := db.QueryRow("SELECT original_url FROM urls WHERE short_id = ?", shortID).Scan(&originalURL)
	if err == sql.ErrNoRows {
		return "", errors.New("URL não encontrada")
	}
	return originalURL, err
}
