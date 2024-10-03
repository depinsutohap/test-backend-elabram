package config

import (
	"context"
	"crypto/tls"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Route() {
	// Create a new HTTP client with a custom transport that disables certificate verification
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient := &http.Client{Transport: transport}

	// Initialize a new Gin router
	router := gin.Default()

	// Set the HTTP client to be used by the Gin server
	router.Use(func(c *gin.Context) {
		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), "httpClient", httpClient))
		c.Next()
	})
	// Create a new Gin router
	r := gin.Default()
	// Define routes
	r.GET("/products", getProducts)
	r.GET("/products/:id", getProduct)
	r.POST("/products", createProduct)
	r.PUT("/products/:id", updateProduct)
	r.DELETE("/products/:id", deleteProduct)

	r.GET("/categories", getCategories)
	r.GET("/categories/:id", getCategory)
	r.POST("/categories", createCategory)
	r.PUT("/categories/:id", updateCategory)
	r.DELETE("/categories/:id", deleteCategory)

	r.GET("/dashboard", dashboardProduct)
	// Start the server
	r.Run(":8080")

	// Add other route groups here

	// Start the Gin server
	err := router.Run(":" + os.Getenv("Port"))
	if err != nil {
		log.Fatalf("Router failed to start: %v\n", err)
	}
}
