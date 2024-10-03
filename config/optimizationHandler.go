package config

import (
	"log"

	"github.com/depinsuthap/elabram-backend-test/entities"
	"github.com/gin-gonic/gin"
)

func retriveAllProduct(c *gin.Context) {
	// Connect to the database
	db := Connect()
	defer db.Close()

	// Execute the query
	rows, err := db.Query("SELECT p.id, p.name, p.description, c.name AS category_name, SUM(oi.quantity) AS total_sold FROM products p JOIN categories c ON p.category_id = c.id JOIN order_items oi ON p.id = oi.product_id GROUP BY p.id, p.name, p.description, c.name")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	type responseProduct struct {
		ID           uint   `json:"id"`
		Name         string `json:"name"`
		Description  string `json:"description"`
		CategoryName string `json:"category_name"`
		TotalSold    uint   `json:"total_sold"`
	}
	// Iterate over the results
	var products []responseProduct
	for rows.Next() {
		var product responseProduct
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.CategoryName, &product.TotalSold)
		if err != nil {
			log.Fatal(err)
		}
		products = append(products, product)
	}
	c.JSON(200, products)
}
func getTopCustomers(c *gin.Context) {
	db := Connect()
	defer db.Close()

	rows, err := db.Query("SELECT c.id, c.name, SUM(o.total) AS total_spent FROM customers c JOIN orders o ON c.id = o.customer_id GROUP BY c.id, c.name ORDER BY total_spent DESC LIMIT 10")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()
	type responseCustomer struct {
		ID         int     `json:"id"`
		Name       string  `json:"name"`
		TotalSpent float64 `json:"total_spent"`
	}
	var customers []responseCustomer

	for rows.Next() {
		var customer responseCustomer
		err := rows.Scan(&customer.ID, &customer.Name, &customer.TotalSpent)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		customers = append(customers, customer)
	}

	c.JSON(200, customers)
}
func getOrderHistory(c *gin.Context) {
	db := Connect()
	defer db.Close()

	customerID := c.Query("customer_id")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	rows, err := db.Query(`SELECT 
                            o.created_at, 
                            c.name AS customer_name, 
                            c.email AS customer_email, 
                            pi.product_id, 
                            p.name AS product_name, 
                            p.price AS product_price, 
                            pi.quantity
                        FROM 
                            orders o
                            JOIN customers c ON o.customer_id = c.id
                            JOIN order_items pi ON o.id = pi.order_id
                            JOIN products p ON pi.product_id = p.id
                        WHERE 
                            o.customer_id = ? 
                            AND o.created_at BETWEEN ? AND ?
                        ORDER BY 
                            o.created_at DESC`, customerID, startDate, endDate)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var orders []entities.Order
	for rows.Next() {
		var order entities.Order
		err := rows.Scan(&order.OrderDate, &order.Customer.Name, &order.Customer.Email, &order.Product.ID, &order.Product.Name, &order.Product.Price, &order.Product.Quantity)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		orders = append(orders, order)
	}

	c.JSON(200, orders)
}
