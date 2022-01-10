package controllers

import (
	"go-react-api/src/database"
	"go-react-api/src/models"

	"github.com/gofiber/fiber/v2"
)

//Orders controller function to retrieve all the orders registered in the system
func Orders(c *fiber.Ctx) error {
	var orders []models.Order
	//preload orderItems as part of order data from DB
	database.DB.Preload("OrderItems").Find(&orders)
	//to replace fullname in every returned row from DB, and the sum of items price
	for i, order := range orders {
		orders[i].Name = order.FullName()
		orders[i].Total = order.GetOrderTotal()

	}
	return c.JSON(orders)
}
