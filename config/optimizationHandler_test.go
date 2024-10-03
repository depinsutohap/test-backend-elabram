// Successfully retrieves all products with their details and total sold quantities
package config_test

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/depinsuthap/elabram-backend-test/config"
	"github.com/depinsuthap/elabram-backend-test/entities"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRetrieveAllProductSuccess(t *testing.T) {
	router := gin.Default()
	router.GET("/products", retriveAllProduct)

	req, _ := http.NewRequest("GET", "/products", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "id")
	assert.Contains(t, w.Body.String(), "name")
	assert.Contains(t, w.Body.String(), "description")
	assert.Contains(t, w.Body.String(), "category_name")
	assert.Contains(t, w.Body.String(), "total_sold")
}
func TestRetrieveAllProductDBConnectionFailure(t *testing.T) {
	originalConnect := config.Connect
	config.Connect = func() (*sql.DB, error) {
		return nil, errors.New("failed to connect to database")
	}
	defer func() { config.Connect = originalConnect }()

	router := gin.Default()
	router.GET("/products", config.retriveAllProduct)

	req, _ := http.NewRequest("GET", "/products", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 500, w.Code)
}

// Retrieves top 10 customers based on total spending
func TestGetTopCustomersSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/top-customers", getTopCustomers)

	req, _ := http.NewRequest(http.MethodGet, "/top-customers", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var customers []responseCustomer
	err := json.Unmarshal(w.Body.Bytes(), &customers)
	assert.NoError(t, err)
	assert.Len(t, customers, 10)
}

// Database connection failure
func TestGetTopCustomersDBConnectionFailure(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Mock the database connection to simulate failure
	originalConnect := Connect
	Connect = func() *sql.DB {
		return nil
	}
	defer func() { Connect = originalConnect }()

	router.GET("/top-customers", getTopCustomers)

	req, _ := http.NewRequest(http.MethodGet, "/top-customers", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response["error"], "connection")
}

// Retrieves order history for a valid customer ID and date range
func TestGetOrderHistoryValidCustomerIDAndDateRange(t *testing.T) {
	router := gin.Default()
	router.GET("/order-history", getOrderHistory)

	req, _ := http.NewRequest("GET", "/order-history?customer_id=123&start_date=2023-01-01&end_date=2023-12-31", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status code 200, got %d", w.Code)
	}

	var orders []entities.Order
	err := json.Unmarshal(w.Body.Bytes(), &orders)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if len(orders) == 0 {
		t.Fatalf("Expected non-empty order history")
	}
}

// Handles missing or invalid customer ID gracefully
func TestGetOrderHistoryMissingOrInvalidCustomerID(t *testing.T) {
	router := gin.Default()
	router.GET("/order-history", getOrderHistory)

	req, _ := http.NewRequest("GET", "/order-history?start_date=2023-01-01&end_date=2023-12-31", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("Expected status code 500, got %d", w.Code)
	}

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if response["error"] == "" {
		t.Fatalf("Expected error message for missing or invalid customer ID")
	}
}
