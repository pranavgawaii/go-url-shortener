package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"go-url-shortener/internal/model"
)

// URLRepository defines the interface for URL storage
type URLRepository interface {
	Create(url *model.URL) error
	GetByShortCode(shortCode string) (*model.URL, error)
}

// postgresURLRepository implements URLRepository for PostgreSQL
type postgresURLRepository struct {
	db *sql.DB
}

// NewURLRepository creates a new instance of URLRepository
func NewURLRepository(db *sql.DB) URLRepository {
	return &postgresURLRepository{db: db}
}

// Create saves a new URL to the database
func (r *postgresURLRepository) Create(url *model.URL) error {
	query := `
		INSERT INTO urls (original_url, short_code, created_at)
		VALUES ($1, $2, $3)
		RETURNING id, click_count
	`
	err := r.db.QueryRow(query, url.OriginalURL, url.ShortCode, url.CreatedAt).
		Scan(&url.ID, &url.ClickCount)
	if err != nil {
		return fmt.Errorf("failed to create url: %w", err)
	}
	return nil
}

// GetByShortCode retrieves a URL by its short code
func (r *postgresURLRepository) GetByShortCode(shortCode string) (*model.URL, error) {
	query := `
		SELECT id, original_url, short_code, click_count, created_at
		FROM urls
		WHERE short_code = $1
	`
	url := &model.URL{}
	err := r.db.QueryRow(query, shortCode).Scan(
		&url.ID,
		&url.OriginalURL,
		&url.ShortCode,
		&url.ClickCount,
		&url.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // Return nil if not found
		}
		return nil, fmt.Errorf("failed to get url: %w", err)
	}
	return url, nil
}
