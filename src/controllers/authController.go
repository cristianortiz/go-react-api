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
	result := database.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}

	return c.JSON(user.Id)
}
func Login(c *fiber.Ctx) error {
	//get data for the request
	var data map[string]string
	//get data from the http request and assign it to data map
	err := c.BodyParser(&data)
	if err != nil {
		return err
	}
	user, founded := UserExists(data["email"])

	if !founded {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"msg": "User or password are wrong",
		})

	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["upassword"]))
	//if pass are differnet return false, the user data does not matter
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"msg": "User or password are wrong",
		})
	}
	return c.JSON(user)

}

//PasswordEncryption encrypts the user pass usin bcrypt library
func PasswordEncryption(pass string) (string, error) {
	//number of layer for encryption algo
	cost := 8
	//GeneratesFormPassword only accepts a slice of bytes []byte
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), cost)
	//return the encrypted password as string
	return string(bytes), err
}

func UserExists(email string) (models.User, bool) {

	var user models.User
	database.DB.Where("email=?", email).First(&user)
	if user.Id == 0 {
		return user, false
	}
	return user, true

}
