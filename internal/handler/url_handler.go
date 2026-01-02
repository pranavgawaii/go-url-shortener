package handler

import (
	"net/http"

	"go-url-shortener/internal/service"

	"github.com/gin-gonic/gin"
)

// URLHandler handles URL related requests
type URLHandler struct {
	service service.URLService
}

// NewURLHandler creates a new URLHandler
func NewURLHandler(service service.URLService) *URLHandler {
	return &URLHandler{service: service}
}

// shortenRequest represents the request body
type shortenRequest struct {
	URL string `json:"url" binding:"required"`
}

// ShortenURL handles the shortening of a URL
func (h *URLHandler) ShortenURL(c *gin.Context) {
	var req shortenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	shortCode, err := h.service.Shorten(req.URL)
	if err != nil {
		// Distinguish errors here if needed (e.g. invalid URL vs internal server error)
		if err.Error() == "invalid URL format" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to shorten URL"})
		return
	}

	// Construct full short URL (using Host header for now)
	// In production, this should likely come from a config variable (BASE_URL)
	shortURL := "http://" + c.Request.Host + "/" + shortCode

	c.JSON(http.StatusCreated, gin.H{
		"short_url": shortURL,
	})
}

// RedirectURL handles the redirection to original URL
func (h *URLHandler) RedirectURL(c *gin.Context) {
	shortCode := c.Param("shortCode")
	if shortCode == "" {
		c.Status(http.StatusBadRequest)
		return
	}

	originalURL, err := h.service.Resolve(shortCode)
	if err != nil {
		if err.Error() == "short code not found" {
			c.Status(http.StatusNotFound)
			return
		}
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Redirect(http.StatusFound, originalURL)
}
