package controllers

import (
	"go-react-api/src/database"
	"go-react-api/src/models"

	"github.com/gofiber/fiber/v2"
)

func GetAmbassadors(c *fiber.Ctx) error {
	var users []models.User

	database.DB.Where("isambassador=true").Find(&users)
	return c.JSON(users)
}
