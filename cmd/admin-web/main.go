package main

import (
	"embed"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

//go:embed static/*
var staticFiles embed.FS

func main() {
	r := gin.Default()

	// Get API URL from environment
	apiURL := getEnv("API_URL", "http://localhost:3000")
	log.Printf("API URL: %s", apiURL)

	// Get embedded static files
	staticFS, err := fs.Sub(staticFiles, "static")
	if err != nil {
		log.Fatalf("Failed to get static files: %v", err)
	}

	// Serve static files
	r.StaticFS("/assets", http.FS(staticFS))

	// Proxy API requests to the actual API server
	r.Any("/api/*path", func(c *gin.Context) {
		proxyRequest(c, apiURL)
	})

	// Proxy admin requests to the actual API server
	r.Any("/admin/*path", func(c *gin.Context) {
		proxyRequest(c, apiURL)
	})

	// Serve index.html for all other routes (SPA)
	r.NoRoute(func(c *gin.Context) {
		data, err := staticFiles.ReadFile("static/index.html")
		if err != nil {
			c.String(http.StatusInternalServerError, "Error loading page")
			return
		}
		c.Data(http.StatusOK, "text/html; charset=utf-8", data)
	})

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "service": "admin-web"})
	})

	// Start server
	port := getEnv("PORT", "3000")
	log.Printf("Admin Web starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func proxyRequest(c *gin.Context, apiURL string) {
	// Build the target URL
	targetURL := apiURL + c.Request.URL.Path
	if c.Request.URL.RawQuery != "" {
		targetURL += "?" + c.Request.URL.RawQuery
	}

	// Create new request
	req, err := http.NewRequest(c.Request.Method, targetURL, c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	// Copy headers
	for key, values := range c.Request.Header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Proxy error: %v", err)
		c.JSON(http.StatusBadGateway, gin.H{"error": "Failed to reach API server"})
		return
	}
	defer resp.Body.Close()

	// Copy response headers
	for key, values := range resp.Header {
		for _, value := range values {
			c.Header(key, value)
		}
	}

	// Copy response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response"})
		return
	}

	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), body)
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
