package config

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/depinsuthap/elabram-backend-test/entities"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/stretchr/testify/assert"
)

type MockRedis struct {
	redis.Conn
}

func NewMockRedis() *MockRedis {
	return &MockRedis{}
}
func (m *MockRedis) Do(command string, args ...interface{}) (interface{}, error) {
	args = m.Called(command, args)
	return args.Get(0), args.Error(1)
}

type RedisConn func() (redis.Conn, error)

// Successfully retrieves categories from Redis and returns them as JSON
func TestGetCategoriesSuccess(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/categories", getCategories)

	// Mock Redis connection
	mockRedis := new(MockRedis)
	mockRedis.On("Do", "KEYS", "category:*").Return([]interface{}{"category:1"}, nil)
	mockRedis.On("Do", "GET", []byte("category:1")).Return([]byte(`{"id":1,"name":"Category1"}`), nil)

	// Perform request
	req, _ := http.NewRequest(http.MethodGet, "/categories", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `[{"id":1,"name":"Category1"}]`, w.Body.String())
}

// Redis connection fails and returns a 500 error
func TestGetCategoriesRedisConnectionFail(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/categories", getCategories)

	// Mock Redis connection failure
	mockRedis := new(MockRedis)
	mockRedis.On("Do", "KEYS", "category:*").Return(nil, errors.New("connection error"))
	RedisConnect = func() RedisConn { return mockRedis }

	// Perform request
	req, _ := http.NewRequest(http.MethodGet, "/categories", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "connection error")
}

// Successfully retrieves category data from Redis cache
func TestGetCategoryFromCache(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/category/:id", getCategory)

	// Mock Redis connection
	redisConn := new(MockRedisConn)
	redisConn.On("Do", "GET", "category:1").Return([]byte(`{"id":1,"name":"Test Category"}`), nil)
	RedisConnect = func() RedisConn { return redisConn }

	// Perform request
	req, _ := http.NewRequest(http.MethodGet, "/category/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)
	var response entities.Category
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, 1, response.ID)
	assert.Equal(t, "Test Category", response.Name)
}

// Handles non-integer category ID gracefully
func TestGetCategoryWithNonIntegerID(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/category/:id", getCategory)

	// Perform request with non-integer ID
	req, _ := http.NewRequest(http.MethodGet, "/category/abc", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusNotFound, w.Code)
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Category not found", response["error"])
}

// Successfully creates a category with valid JSON input
func TestCreateCategoryWithValidJSON(t *testing.T) {
	router := gin.Default()
	router.POST("/category", createCategory)

	validJSON := `{"name": "New Category", "description": "A new category description"}`
	req, _ := http.NewRequest("POST", "/category", strings.NewReader(validJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, but got %d", http.StatusCreated, w.Code)
	}

	var category entities.Category
	err := json.Unmarshal(w.Body.Bytes(), &category)
	if err != nil {
		t.Errorf("Error unmarshalling response: %v", err)
	}

	if category.Name != "New Category" || category.Description != "A new category description" {
		t.Errorf("Category data does not match input")
	}
}

// Handles invalid JSON input gracefully with HTTP 400 error
func TestCreateCategoryWithInvalidJSON(t *testing.T) {
	router := gin.Default()
	router.POST("/category", createCategory)

	invalidJSON := `{"name": "New Category", "description": }`
	req, _ := http.NewRequest("POST", "/category", strings.NewReader(invalidJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, w.Code)
	}

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Error unmarshalling response: %v", err)
	}

	if response["error"] == "" {
		t.Errorf("Expected error message in response")
	}
}

// Successfully updates a category with valid ID and JSON data
func TestUpdateCategorySuccess(t *testing.T) {
	router := gin.Default()
	router.PUT("/category/:id", updateCategory)

	mockDB := new(MockDB)
	mockDB.On("Exec", "UPDATE categories SET name = ?, description = ? WHERE id = ?", "New Name", "New Description", 1).Return(nil, nil)
	Connect = func() *sql.DB { return mockDB }

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/category/1", strings.NewReader(`{"name": "New Name", "description": "New Description"}`))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var response entities.Category
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "New Name", response.Name)
	assert.Equal(t, "New Description", response.Description)
}

// Handles non-integer ID parameter gracefully
func TestUpdateCategoryNonIntegerID(t *testing.T) {
	router := gin.Default()
	router.PUT("/category/:id", updateCategory)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/category/abc", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "Category not found", response["error"])
}

// Successfully delete a category with a valid ID
func TestDeleteCategoryWithValidID(t *testing.T) {
	router := gin.Default()
	router.DELETE("/category/:id", deleteCategory)

	req, _ := http.NewRequest("DELETE", "/category/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}

	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)

	if response["message"] != "Category deleted successfully" {
		t.Errorf("Expected message 'Category deleted successfully', got %s", response["message"])
	}
}

// Handle non-integer ID parameter gracefully
func TestDeleteCategoryWithNonIntegerID(t *testing.T) {
	router := gin.Default()
	router.DELETE("/category/:id", deleteCategory)

	req, _ := http.NewRequest("DELETE", "/category/abc", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status code 404, got %d", w.Code)
	}

	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)

	if response["error"] != "Category not found" {
		t.Errorf("Expected error 'Category not found', got %s", response["error"])
	}
}
