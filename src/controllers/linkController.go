package controllers

import (
	"go-react-api/src/database"
	"go-react-api/src/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

//Link controller function to retrieve all the links of a specific user
func Link(c *fiber.Ctx) error {
	//user id in http request is a string, cast to the int type for query to DB
	id, _ := strconv.Atoi(c.Params("id"))

	var links []models.Link

	database.DB.Where("user_id = ?", id).Find(&links)
	for i, link := range links {
		var orders []models.Order
		database.DB.Where("code=? and complete=true", link.Code).Find(&orders)
		links[i].Orders = orders
	}
	return c.JSON(links)

}
