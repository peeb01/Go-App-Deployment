package api

import (
    "github.com/gofiber/fiber/v2"
)

func Helloworld(c *fiber.Ctx) error {
    return c.JSON(fiber.Map{"message": "Hello"})
}