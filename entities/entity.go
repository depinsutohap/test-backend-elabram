package entities

// Product represents a product
type Product struct {
	ID            int     `db:"id" json:"id"`
	Name          string  `db:"name" json:"name"`
	Description   string  `db:"description" json:"description"`
	Price         float64 `db:"price" json:"price"`
	CategoryID    int     `db:"category_id" json:"category_id"`
	StockQuantity int     `db:"stock_quantity" json:"stock_quantity"`
	IsActive      bool    `db:"is_active" json:"is_active"`
	CreatedAt     string  `db:"created_at" json:"created_at"`
	UpdatedAt     string  `db:"updated_at" json:"updated_at"`
}

// Category represents a category
type Category struct {
	ID          int    `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
}

type ProductDasboard struct {
	ID            int      `db:"id" json:"id"`
	Name          string   `db:"name" json:"name"`
	Description   string   `db:"description" json:"description"`
	Price         float64  `db:"price" json:"price"`
	CategoryID    int      `db:"category_id" json:"category_id"`
	StockQuantity int      `db:"stock_quantity" json:"stock_quantity"`
	IsActive      bool     `db:"is_active" json:"is_active"`
	CreatedAt     string   `db:"created_at" json:"created_at"`
	UpdatedAt     string   `db:"updated_at" json:"updated_at"`
	Category      Category `db:"category" json:"category"`
}
type Order struct {
	OrderDate string       `json:"order_date"`
	Customer  Customer     `json:"customer"`
	Product   ProductOrder `json:"products"`
}

// Product represents a product
type ProductOrder struct {
	ID          int     `db:"id" json:"id"`
	Name        string  `db:"name" json:"name"`
	Description string  `db:"description" json:"description"`
	Price       float64 `db:"price" json:"price"`
	CategoryID  int     `db:"category_id" json:"category_id"`
	Quantity    int     `db:"quantity" json:"quantity"`
}
type Customer struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
