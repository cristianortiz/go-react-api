package controllers

import (
	"go-react-api/src/database"
	"go-react-api/src/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

//GetProducts is the controller func response the request to get all the products from DB
func GetProducts(c *fiber.Ctx) error {
	var products []models.Product

	database.DB.Find(&products)
	return c.JSON(products)
}

//CreateProduct id the func to inset a new product in DB
func CreateProduct(c *fiber.Ctx) error {
	//product model type to store the data from the http request
	var product models.Product
	//get data from the http request and assign it to products models type
	err := c.BodyParser(&product)
	if err != nil {
		return err
	}
	//insert new product data in DB
	//insert new user data un Db
	result := database.DB.Create(&product)
	if result.Error != nil {
		return result.Error
	}

	return c.JSON(product)
}

func GetProductByID(c *fiber.Ctx) error {
	var product models.Product

	id, _ := strconv.Atoi(c.Params("id"))

	product.Id = uint(id)
	result := database.DB.Where("id = ?", id).First(&product)
	if result.Error != nil {
		return result.Error
	}

	return c.JSON(product)

}
