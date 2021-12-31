package controllers

import (
	"go-react-api/src/database"
	"go-react-api/src/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

//GetProducts is the controller func to response the request to get all the products from DB
func GetProducts(c *fiber.Ctx) error {
	var products []models.Product

	database.DB.Find(&products)
	return c.JSON(products)
}

//CreateProduct is the func to insert a new product in DB
func CreateProduct(c *fiber.Ctx) error {
	//product model type to store the data from the http request
	var product models.Product
	//get data from the http request and assign it to products models type
	err := c.BodyParser(&product)
	if err != nil {
		return err
	}
	//insert new product data in DB
	result := database.DB.Create(&product)
	if result.Error != nil {
		return result.Error
	}

	return c.JSON(product)
}

//GetProductByID, controller function to retrieve the info a product with their id
func GetProductByID(c *fiber.Ctx) error {
	var product models.Product
	//id in http request is a string, cast to the int type for query to DB
	id, _ := strconv.Atoi(c.Params("id"))

	product.Id = uint(id)
	result := database.DB.Where("id = ?", id).First(&product)
	if result.Error != nil {
		return result.Error
	}

	return c.JSON(product)

}

//UpdateProduct, controller function to update the info a product with their id
func UpdateProduct(c *fiber.Ctx) error {
	//id in http request is a string, cast to the int type for query to DB
	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{
		Id: uint(id),
	}
	err := c.BodyParser(&product)
	if err != nil {
		return err
	}

	result := database.DB.Model(&product).Updates(&product)
	if result.Error != nil {
		return result.Error
	}

	return c.JSON(product)

}
func DeleteProduct(c *fiber.Ctx) error {
	//id in http request is a string, cast to the int type for query to DB
	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{
		Id: uint(id),
	}
	err := c.BodyParser(&product)
	if err != nil {
		return err
	}

	result := database.DB.Delete(&product)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
