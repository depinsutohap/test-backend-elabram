package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/depinsuthap/elabram-backend-test/entities"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Retrieve all products from Redis cache successfully
func TestGetProductsSuccess(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/products", getProducts)

	// Mock Redis connection and data
	mockRedis := new(MockRedis)
	mockRedis.On("Do", "KEYS", "product:*").Return([]interface{}{"product:1", "product:2"}, nil)
	mockRedis.On("Do", "GET", []byte("product:1")).Return([]byte(`{"id":1,"name":"Product 1"}`), nil)
	mockRedis.On("Do", "GET", []byte("product:2")).Return([]byte(`{"id":2,"name":"Product 2"}`), nil)
	RedisConnect = func() RedisConn { return mockRedis }

	// Perform request
	req, _ := http.NewRequest(http.MethodGet, "/products", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)
	var products []entities.Product
	err := json.Unmarshal(w.Body.Bytes(), &products)
	assert.NoError(t, err)
	assert.Len(t, products, 2)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 2", products[1].Name)
}

// Redis connection fails
func TestGetProductsRedisConnectionFail(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/products", getProducts)

	// Mock Redis connection failure
	mockRedis := new(MockRedis)
	mockRedis.On("Do", "KEYS", "product:*").Return(nil, fmt.Errorf("connection error"))
	RedisConnect = func() RedisConn { return mockRedis }

	// Perform request
	req, _ := http.NewRequest(http.MethodGet, "/products", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "connection error", response["error"])
}

// Successfully creates a product when valid JSON is provided
func TestCreateProductSuccess(t *testing.T) {
	router := gin.Default()
	router.POST("/product", createProduct)

	productJSON := `{
			"name": "Test Product",
			"description": "A product for testing",
			"price": 99.99,
			"category_id": 1,
			"stock_quantity": 10,
			"is_active": true
		}`

	req, _ := http.NewRequest("POST", "/product", strings.NewReader(productJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, but got %d", http.StatusCreated, w.Code)
	}

	var product entities.Product
	err := json.Unmarshal(w.Body.Bytes(), &product)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if product.Name != "Test Product" {
		t.Errorf("Expected product name 'Test Product', but got '%s'", product.Name)
	}
}

// Handles invalid JSON input gracefully by returning a 400 error
func TestCreateProductInvalidJSON(t *testing.T) {
	router := gin.Default()
	router.POST("/product", createProduct)

	invalidJSON := `{
			"name": "Test Product",
			"description": "A product for testing",
			"price": "invalid_price"
		}`

	req, _ := http.NewRequest("POST", "/product", strings.NewReader(invalidJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, w.Code)
	}

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if response["error"] == "" {
		t.Error("Expected an error message in the response")
	}
}

// Successfully updates a product when valid data is provided
func TestUpdateProductSuccess(t *testing.T) {
	router := gin.Default()
	router.PUT("/product/:id", updateProduct)

	product := entities.Product{
		Name:          "Updated Product",
		Description:   "Updated Description",
		Price:         99.99,
		CategoryID:    1,
		StockQuantity: 10,
		IsActive:      true,
	}

	jsonData, _ := json.Marshal(product)
	req, _ := http.NewRequest("PUT", "/product/1", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var updatedProduct entities.Product
	json.Unmarshal(w.Body.Bytes(), &updatedProduct)
	assert.Equal(t, product.Name, updatedProduct.Name)
	assert.Equal(t, product.Description, updatedProduct.Description)
}

// Handles non-integer product ID gracefully by returning a 404 error
func TestUpdateProductNonIntegerID(t *testing.T) {
	router := gin.Default()
	router.PUT("/product/:id", updateProduct)

	req, _ := http.NewRequest("PUT", "/product/abc", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "Product not found", response["error"])
}

// Successfully delete a product with a valid ID
func TestDeleteProductWithValidID(t *testing.T) {
	router := gin.Default()
	router.DELETE("/product/:id", deleteProduct)

	req, _ := http.NewRequest("DELETE", "/product/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}

	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)

	if response["message"] != "Product deleted successfully" {
		t.Errorf("Expected message 'Product deleted successfully', got '%s'", response["message"])
	}
}

// Handle non-numeric ID input gracefully
func TestDeleteProductWithNonNumericID(t *testing.T) {
	router := gin.Default()
	router.DELETE("/product/:id", deleteProduct)

	req, _ := http.NewRequest("DELETE", "/product/abc", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status code 404, got %d", w.Code)
	}

	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)

	if response["error"] != "Product not found" {
		t.Errorf("Expected error 'Product not found', got '%s'", response["error"])
	}
}

// Correctly parses 'limit' and 'page' query parameters
func TestCorrectParsingOfLimitAndPageParameters(t *testing.T) {
	router := gin.Default()
	router.GET("/dashboard", dashboardProduct)

	req, _ := http.NewRequest("GET", "/dashboard?limit=20&page=2", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status code 200, got %d", w.Code)
	}

	var response []entities.ProductDasboard
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}

	// Additional assertions can be added here to verify the response content
}

// Handles non-integer 'limit' and 'page' query parameters gracefully
func TestNonIntegerLimitAndPageParameters(t *testing.T) {
	router := gin.Default()
	router.GET("/dashboard", dashboardProduct)

	req, _ := http.NewRequest("GET", "/dashboard?limit=abc&page=xyz", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status code 200, got %d", w.Code)
	}

	var response []entities.ProductDasboard
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}

	// Additional assertions can be added here to verify the response content
}
