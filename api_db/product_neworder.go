package api_db

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"fmt"
	"time"
)

func timePtr(t time.Time) *time.Time {
	return &t
}


func NewOrder(db *gorm.DB, c *fiber.Ctx) error {
	var request CreateOrderRequest

	// parser JSON request body
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	// Find customer by ID
	var customer Customer
	if err := db.First(&customer, request.CustomerID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Customer not found",
		})
	}

	// Create an order record
	order := Order{
		CustomerID: request.CustomerID,
		OrderDate:  timePtr(time.Now()), // current time
		Status:     1,
		TotalPrice: 0,
	}
	if err := db.Create(&order).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not create order",
		})
	}

	// init total price
	totalPrice := float32(0)

	// Process each product in the order
	for _, item := range request.Products {
		var product Product
		if err := db.First(&product, item.ProductID).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": fmt.Sprintf("Product with ID %d not found", item.ProductID),
			})
		}

		// Check stock availability
		if product.StockQuantity < item.Quantity {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": fmt.Sprintf("Not enough stock for product ID %d", item.ProductID),
			})
		}

		// Calculate price
		priceAtOrder := product.Price * float32(item.Quantity)
		totalPrice += priceAtOrder

		// Create an order detail
		orderDetail := OrderDetail{
			OrderID:      int(order.ID),
			ProductID:    item.ProductID,
			Quantity:     item.Quantity,
			PriceAtOrder: product.Price,
		}
		if err := db.Create(&orderDetail).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Could not create order detail",
			})
		}

		// Update product stock
		product.StockQuantity -= item.Quantity
		if err := db.Save(&product).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Could not update product stock",
			})
		}
	}

	// Update the order's total price
	order.TotalPrice = totalPrice
	if err := db.Save(&order).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not update order total price",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(order)
}
