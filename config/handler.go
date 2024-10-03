package config

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/depinsuthap/elabram-backend-test/entities"
	"github.com/gin-gonic/gin"
)

func getProducts(c *gin.Context) {
	// Check if the products data is cached
	var data []entities.Product
	redisConn := RedisConnect()
	defer redisConn.Close()
	keys, err := redisConn.Do("KEYS", "product:*")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	for _, k := range keys.([]interface{}) {
		var product entities.Product
		reply, err := redisConn.Do("GET", k.([]byte))

		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		if err := json.Unmarshal(reply.([]byte), &product); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		data = append(data, product)
	}
	c.JSON(200, data)
}

func getProduct(c *gin.Context) {
	var data entities.Product
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(404, gin.H{"error": "Product not found"})
		return
	}
	redisConn := RedisConnect()
	defer redisConn.Close()
	// Check if the category data is cached
	reply, err := redisConn.Do("GET", "product:"+strconv.Itoa(id))
	if err == nil && reply != nil {
		if err := json.Unmarshal(reply.([]byte), &data); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	} else if reply == nil {
		c.JSON(404, gin.H{"error": "Product not found"})
		return

	}
	c.JSON(200, data)
}
func createProduct(c *gin.Context) {
	var product entities.Product
	err := c.BindJSON(&product)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	db := Connect()
	defer db.Close()
	_, err = db.Exec("INSERT INTO products (name, description, price, category_id, stock_quantity, is_active, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, NOW(), NOW())", product.Name, product.Description, product.Price, product.CategoryID, product.StockQuantity, product.IsActive)
	db.QueryRow("SELECT LAST_INSERT_ID();").Scan(&product.ID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}
	b, _ := json.Marshal(product)
	RedisCmd(b, "product:"+strconv.Itoa(product.ID), "SET")
	c.JSON(201, product)
}
func updateProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(404, gin.H{"error": "Product not found"})
		return
	}
	var product entities.Product
	err = c.BindJSON(&product)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	product.ID = id
	db := Connect()
	defer db.Close()
	_, err = db.Exec("UPDATE products SET name = ?, description = ?, price = ?, category_id = ?, stock_quantity = ?, is_active = ? WHERE id = ?", product.Name, product.Description, product.Price, product.CategoryID, product.StockQuantity, product.IsActive, product.ID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	b, _ := json.Marshal(product)
	RedisCmd(b, "product:"+strconv.Itoa(product.ID), "SET")
	c.JSON(200, product)
}
func deleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(404, gin.H{"error": "Product not found"})
		return
	}
	db := Connect()
	defer db.Close()
	result, err := db.Exec("DELETE FROM products WHERE id = ?", id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	affected, _ := result.RowsAffected()
	fmt.Println(affected)
	if affected > 0 {
		RedisCmd([]byte{}, "product:"+strconv.Itoa(id), "DEL")
	}
	c.JSON(200, gin.H{"message": "Product deleted successfully"})
}
func getCategories(c *gin.Context) {
	// Check if the categories data is cached
	var data []entities.Category
	redisConn := RedisConnect()
	defer redisConn.Close()
	keys, err := redisConn.Do("KEYS", "category:*")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	for _, k := range keys.([]interface{}) {
		var category entities.Category
		reply, err := redisConn.Do("GET", k.([]byte))

		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		if err := json.Unmarshal(reply.([]byte), &category); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
		}
		data = append(data, category)
	}
	c.JSON(200, data)
}
func getCategory(c *gin.Context) {
	var data entities.Category
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(404, gin.H{"error": "Category not found"})
		return
	}
	redisConn := RedisConnect()
	defer redisConn.Close()
	// Check if the category data is cached
	reply, err := redisConn.Do("GET", "category:"+strconv.Itoa(id))
	if err == nil && reply != nil {
		if err := json.Unmarshal(reply.([]byte), &data); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	} else if reply == nil {
		c.JSON(404, gin.H{"error": "Category not found"})
		return

	}
	c.JSON(200, data)
}
func createCategory(c *gin.Context) {
	var category entities.Category
	err := c.BindJSON(&category)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	db := Connect()
	defer db.Close()
	_, err = db.Exec("INSERT INTO categories (name, description) VALUES (?, ?);", category.Name, category.Description)
	db.QueryRow("SELECT LAST_INSERT_ID();").Scan(&category.ID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	b, err := json.Marshal(category)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	RedisCmd(b, "category:"+strconv.Itoa(category.ID), "SET")
	c.JSON(201, category)
}
func updateCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(404, gin.H{"error": "Category not found"})
		return
	}
	var category entities.Category
	err = c.BindJSON(&category)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	category.ID = id
	db := Connect()
	defer db.Close()
	_, err = db.Exec("UPDATE categories SET name = ?, description = ? WHERE id = ?", category.Name, category.Description, category.ID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	b, err := json.Marshal(category)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	RedisCmd(b, "category:"+strconv.Itoa(category.ID), "SET")
	c.JSON(200, category)
}
func deleteCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(404, gin.H{"error": "Category not found"})
		return
	}
	db := Connect()
	defer db.Close()
	result, err := db.Exec("DELETE FROM categories WHERE id = ?", id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	affected, _ := result.RowsAffected()
	if affected > 0 {
		RedisCmd([]byte{}, "category:"+strconv.Itoa(id), "DEL")
	}
	c.JSON(200, gin.H{"message": "Category deleted successfully"})
}
func dashboardProduct(c *gin.Context) {
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 10
	}
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 0
	}
	if page > 0 {
		page = page * limit
	}
	sortType := strings.ToLower(c.Query("sort_type"))
	switch sortType {
	case "asc":
		sortType = "ASC"
	case "desc":
		sortType = "DESC"
	default:
		sortType = "ASC"
	}
	sortColumn := strings.ToLower(c.Query("sort_column"))
	switch sortColumn {
	case "name":
		sortColumn = "p.name"
	case "price":
		sortColumn = "p.price"
	case "category":
		sortColumn = "c.name"
	case "stock_quantity":
		sortColumn = "p.stock_quantity"
	default:
		sortColumn = "p.name"
	}
	name := c.Query("name")
	categoryId := c.Query("category_id")
	if categoryId != "" {
		categoryId = "p.category_id in (" + categoryId + ")"
	} else {
		categoryId = "1"
	}
	priceMin := c.Query("price_min")
	if regexp.MustCompile(`\d`).MatchString(priceMin) && priceMin != "" {
		priceMin = "p.price >= " + priceMin
	} else {
		priceMin = "1"
	}
	priceMax := c.Query("price_max")
	if regexp.MustCompile(`\d`).MatchString(priceMax) && priceMax != "" {
		priceMax = "p.price <= " + priceMax
	} else {
		priceMax = "1"
	}
	stockMin := c.Query("stock_quantity_min")
	if regexp.MustCompile(`\d`).MatchString(stockMin) && stockMin != "" {
		stockMin = "p.price >= " + stockMin
	} else {
		stockMin = "1"
	}
	stockMax := c.Query("stock_quantity_max")
	if regexp.MustCompile(`\d`).MatchString(stockMax) && stockMax != "" {
		stockMax = "p.price <= " + stockMax
	} else {
		stockMax = "1"
	}
	var data []entities.ProductDasboard
	db := Connect()
	defer db.Close()
	rows, err := db.Query("SELECT p.id FROM products p WHERE p.name LIKE '%" + name + "%' AND " + categoryId + " AND " + priceMin + " AND " + priceMax + " AND " + stockMin + " AND " + stockMax + " ORDER BY " + sortColumn + " " + sortType + " LIMIT " + strconv.Itoa(page) + "," + strconv.Itoa(limit))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	redisConn := RedisConnect()
	defer redisConn.Close()
	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		var product entities.ProductDasboard
		reply, _ := redisConn.Do("GET", "product:"+strconv.Itoa(id))
		if err := json.Unmarshal(reply.([]byte), &product); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		reply, _ = redisConn.Do("GET", "category:"+strconv.Itoa(product.CategoryID))
		if err := json.Unmarshal(reply.([]byte), &product.Category); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		data = append(data, product)
	}
	c.JSON(200, data)
}
