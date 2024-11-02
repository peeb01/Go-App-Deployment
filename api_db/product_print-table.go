package api_db

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)


func GetCustomer(db *gorm.DB, c *fiber.Ctx) error {
	var cnsmr []Customer
	if result := db.Find(&cnsmr); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch Customer",
		})
	}
	return c.Status(fiber.StatusOK).JSON(cnsmr)
}

func GetProduct(db *gorm.DB, c *fiber.Ctx) error {
	var prod []Product
	if result := db.Find(&prod); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch Product",
		})
	}
	return c.Status(fiber.StatusOK).JSON(prod)
}

func GetOrder(db *gorm.DB, c *fiber.Ctx) error {
	var orders []Order
	if result := db.Find(&orders); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Fail to fetch Order",
		})
	}
	return c.Status(fiber.StatusOK).JSON(orders)
}

func GetOrderDetail(db *gorm.DB, c *fiber.Ctx) error {
	var detail []OrderDetail
	if result := db.Find(&detail); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch Order Detail",
		})
	}
	return c.Status(fiber.StatusOK).JSON(detail)
}
