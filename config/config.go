package config

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/depinsuthap/elabram-backend-test/entities"
	"github.com/gomodule/redigo/redis"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3307)/database_1")
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Connection established")
	return db
}
func initDB(db sql.DB) {
	// cached data to redist
	rows, err := db.Query("SELECT * FROM categories")
	if err != nil {
		log.Println(err)
		return
	}
	redisConn := RedisConnect()
	defer rows.Close()
	defer redisConn.Close()
	keys, err := redisConn.Do("KEYS", "category:*")
	if err != nil {
		panic(err.Error())
	}
	for _, k := range keys.([]interface{}) {
		_, err := redisConn.Do("DEL", k.([]byte))

		if err != nil {
			panic(err.Error())
		}
	}
	for rows.Next() {
		var category entities.Category
		err = rows.Scan(&category.ID, &category.Name, &category.Description)
		if err != nil {
			panic(err.Error())
		}
		b, err := json.Marshal(category)
		if err != nil {
			panic(err.Error())
		}
		_, err = redisConn.Do("SET", "category:"+strconv.Itoa(category.ID), b)
		if err != nil {
			panic(err.Error())
		}
	}
	fmt.Println("Init Data Category Done .....")

	keys, err = redisConn.Do("KEYS", "product:*")
	if err != nil {
		panic(err.Error())
	}
	for _, k := range keys.([]interface{}) {
		_, err := redisConn.Do("DEL", k.([]byte))

		if err != nil {
			panic(err.Error())
		}
	}
	rows, err = db.Query("SELECT * FROM products")
	if err != nil {
		log.Println(err)
		return
	}
	for rows.Next() {
		var product entities.Product
		err = rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CategoryID, &product.StockQuantity, &product.IsActive, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			panic(err.Error())
		}
		b, err := json.Marshal(product)
		if err != nil {
			panic(err.Error())
		}
		_, err = redisConn.Do("SET", "product:"+strconv.Itoa(product.ID), b)
		if err != nil {
			panic(err.Error())
		}
	}

	fmt.Println("Init Data Products Done .....")
}

func RedisCmd(b []byte, key string, cmd string) {
	redisConn := RedisConnect()
	defer redisConn.Close()
	switch cmd {
	case "SET":
		_, err := redisConn.Do(cmd, key, b)
		if err != nil {
			panic(err.Error())
		}
	case "DEL":
		result, err := redisConn.Do(cmd, key)
		fmt.Println(result)
		if err != nil {
			panic(err.Error())
		}
	}
}

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}

func RedisConnect() redis.Conn {
	redisConn, err := redis.Dial("tcp", "localhost:6379")
	HandleError(err)
	return redisConn
}
