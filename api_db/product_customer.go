package api_db

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

)


func RegisterCustomer(db *gorm.DB, c *fiber.Ctx) error {
	var request Customer

	// Parser the request body
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	// Check for required fields
	if request.CustomerName == "" || request.Email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Name and email are required fields",
		})
	}

	// Check if the email is already registered
	var existingCustomer Customer
	if err := db.Where("email = ?", request.Email).First(&existingCustomer).Error; err == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "Email is already registered",
		})
	}

	// Create a new customer record
	newCustomer := Customer{
		CustomerName: request.CustomerName,
		Address:      request.Address,
		Phone:        request.Phone,
		Email:        request.Email,
	}

	if err := db.Create(&newCustomer).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not create customer",
		})
	}

	// Return the created customer as JSON, excluding sensitive fields like ID and created_at
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id":      newCustomer.ID,
		"name":    newCustomer.CustomerName,
		"address": newCustomer.Address,
		"phone":   newCustomer.Phone,
		"email":   newCustomer.Email,
	})
}
