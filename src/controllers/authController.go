package controllers

import (
	"go-react-api/src/database"
	"go-react-api/src/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	//get data for the request
	var data map[string]string
	//get data from the http request and assign it to data map
	err := c.BodyParser(&data)
	if err != nil {
		return err
	}
	//validations
	if data["upassword"] != data["upassword_confirm"] {
		return c.JSON(fiber.Map{
			"message": "passwords do not match",
		})
	}
	password, err := PasswordEncryption(data["upassword"])
	if err != nil {
		return err
	}
	//in this case the register is for an admin, so isAmbassador is false by default
	user := models.User{
		FirstName:    data["firstname"],
		LastName:     data["lastname"],
		Email:        data["email"],
		Password:     password,
		IsAmbassador: false,
	}
	//insert new user data un Db
	database.DB.Create(&user)

	return c.JSON(fiber.Map{
		"message": "hello again",
	})
}

//PasswordEncryption encrypts the user pass usin bcrypt library
func PasswordEncryption(pass string) (string, error) {
	//number of layer for encryption algo
	cost := 8
	//GeneratesFormPassword only accepts a slice of bytes []byte
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), cost)
	return string(bytes), err
}
