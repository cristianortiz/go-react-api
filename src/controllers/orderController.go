package controllers

import (
	"go-react-api/src/database"
	"go-react-api/src/models"

	"github.com/gofiber/fiber/v2"
)

//Orders controller function to retrieve all the orders registered in the system
func Orders(c *fiber.Ctx) error {
	var orders []models.Order

	database.DB.Find(&orders)
	return c.JSON(orders)
}
