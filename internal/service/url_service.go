package service

import (
	"crypto/rand"
	"errors"
	"math/big"
	"net/url"
	"strings"
	"time"

	"go-url-shortener/internal/model"
	"go-url-shortener/internal/repository"
)

const (
	base62Chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	codeLength  = 6
)

var (
	ErrOriginalURLRequired = errors.New("original URL is required")
	ErrInvalidURLFormat    = errors.New("invalid URL format")
	ErrShortCodeNotFound   = errors.New("short code not found")
)

// URLService defines the interface for URL business logic
type URLService interface {
	Shorten(originalURL string) (string, error)
	Resolve(shortCode string) (string, error)
}

type urlService struct {
	repo repository.URLRepository
}

// NewURLService creates a new instance of URLService
func NewURLService(repo repository.URLRepository) URLService {
	return &urlService{repo: repo}
}

// Shorten validates a URL, generates a short code, and saves it
func (s *urlService) Shorten(originalURL string) (string, error) {
	// 1. Basic Validation
	if originalURL == "" {
		return "", ErrOriginalURLRequired
	}
	// Ensure scheme exists
	if !strings.HasPrefix(originalURL, "http://") && !strings.HasPrefix(originalURL, "https://") {
		originalURL = "http://" + originalURL
	}
	// Parse URL to check validity
	if _, err := url.ParseRequestURI(originalURL); err != nil {
		return "", ErrInvalidURLFormat
	}

	// 2. Generate and Save with Retry
	maxRetries := 5
	for i := 0; i < maxRetries; i++ {
		code, err := generateShortCode(codeLength)
		if err != nil {
			return "", err
		}

		// Check for collision
		existing, err := s.repo.GetByShortCode(code)
		if err != nil {
			return "", err // DB error
		}
		if existing != nil {
			continue // Collision, retry
		}

		// Save to DB
		newURL := &model.URL{
			OriginalURL: originalURL,
			ShortCode:   code,
			CreatedAt:   time.Now(),
		}

		if err := s.repo.Create(newURL); err != nil {
			return "", err // Could be race condition, return error for now
		}

		return code, nil
	}

	return "", errors.New("failed to generate unique short code after retries")
}

// Resolve retrieves the original URL and increments click count
func (s *urlService) Resolve(shortCode string) (string, error) {
	url, err := s.repo.GetByShortCode(shortCode)
	if err != nil {
		return "", err
	}
	if url == nil {
		return "", ErrShortCodeNotFound
	}

	// Increment click count asynchronously (fire and forget, or handle error if critical)
	// For now, we'll do it synchronously for simplicity
	if err := s.repo.IncrementClickCount(shortCode); err != nil {
		// Log error but don't fail the redirect?
		// log.Printf("Failed to increment click count: %v", err)
	}

	return url.OriginalURL, nil
}

// generateShortCode generates a random base62 string
func generateShortCode(length int) (string, error) {
	bytes := make([]byte, length)
	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(base62Chars))))
		if err != nil {
			return "", err
		}
		bytes[i] = base62Chars[num.Int64()]
	}
	return string(bytes), nil
}
