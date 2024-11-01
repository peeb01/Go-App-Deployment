package api_db

import (
	"gorm.io/gorm"
	"time"
)

type Customer struct {
	gorm.Model
	CustomerName	string `json:"name"`
	Address  		string `json:"address"`
	Phone			string `json:"phone"`
	Email			string `json:"email"`
}

type Product struct {
	gorm.Model
	ProductName		string `json:"name"`
	Description		string `json:"description"`
	Price			float32 `json:"price"`
	StockQuantity	int `json:"stockquantity"`
}

type Order struct {
	gorm.Model
	CustomerID		int       `json:"customer_id"` // Foreign key
	Customer		Customer  `gorm:"foreignKey:CustomerID"` // Relationship
	OrderDate		*time.Time
	TotalPrice		float32   `json:"total_price"`
	Status			int       `json:"status"`
}

type OrderDetail struct {
	gorm.Model
	OrderID  		int      `json:"order_id"`  // Foreign key
	Order			Order    `gorm:"foreignKey:OrderID"` // Relationship
	ProductID		int      `json:"product_id"` // Foreign key
	Product			Product  `gorm:"foreignKey:ProductID"` // Relationship
	Quantity		int      `json:"quantity"`
	PriceAtOrder	float32  `json:"priceatorder"`
}

type CreateOrderRequest struct {
	CustomerID int               `json:"customer_id"` // Foreign key to Customer
	Products   []OrderProduct    `json:"products"`    // List of products being ordered
}

type OrderProduct struct {
	ProductID int `json:"product_id"` // Foreign key to Product
	Quantity  int `json:"quantity"`   // Number of items ordered for the product
}
