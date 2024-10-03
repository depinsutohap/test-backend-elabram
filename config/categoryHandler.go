package config

import (
	"encoding/json"
	"strconv"

	"github.com/depinsuthap/elabram-backend-test/entities"
	"github.com/gin-gonic/gin"
)

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
