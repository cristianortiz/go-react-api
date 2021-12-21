package controllers

import (
	"go-react-api/src/database"
	"go-react-api/src/models"
	"strconv"

	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

//Register creates a new user  in DB with an encrypted password
func Register(c *fiber.Ctx) error {
	//map to store the request data
	var data map[string]string
	//get data from the http request and assign it to data map
	err := c.BodyParser(&data)
	if err != nil {
		return err
	}
	//validations
	//check if the email is already in use
	_, used := UserEmailExists(data["email"])
	if used {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"msg": "Username is alreay in use",
		})
	}
	if data["upassword"] != data["upassword_confirm"] {
		return c.JSON(fiber.Map{
			"message": "passwords do not match",
		})
	}

	//in this case the register is for an admin, so isAmbassador is false by default
	user := models.User{
		FirstName:    data["firstname"],
		LastName:     data["lastname"],
		Email:        data["email"],
		IsAmbassador: false,
	}
	//password encryption
	user.SetAndEncryptPassword(data["upassword"])
	//insert new user data un Db
	result := database.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}

	return c.JSON(user.Id)
}

//Login function checks if an user exists and verify their password
func Login(c *fiber.Ctx) error {
	//map to store the request data
	var data map[string]string
	//get data from the http request and assign it to data map
	err := c.BodyParser(&data)
	if err != nil {
		return err
	}
	//check if the email is already in use
	user, founded := UserEmailExists(data["email"])

	if !founded {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"msg": "Username or password are wrong",
		})
	}
	//bcrypt only works with slice of bytes data,hash the password received as parameter
	//and the pass returned by the DB
	err = user.ComparePassword(data["upassword"])
	//if pass are not equals  response the error
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"msg": "Username or password are wrong",
		})
	}
	//generates JWT for use auth
	token, err := GeneratesJWT(user)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"msg": "Invalid Credentials",
		})
	}
	//the JWT recorded in a cookie
	//cookie expiration time
	expirationTime := time.Now().Add(24 * time.Hour)
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  expirationTime,
		HTTPOnly: true, //to be sent to the backend
	}
	//set cookie in fiber context
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"msg:": "sucsess",
	})
}

//AuthenticatedUser returns the data of a logged user, using the jwt in cookie stored in fiber context
func AuthenticatedUser(c *fiber.Ctx) error {
	//get the cookie from fiber context
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("ReactGoAPI"), nil
	})
	if err != nil || !token.Valid {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"msg": "Invalid credentials",
		})
	}
	payload := token.Claims.(*jwt.StandardClaims)
	//query the user data from DB with their id stored in jwt
	var user models.User
	//payload.Subject stores the id of the user logged
	database.DB.Where("id = ?", payload.Subject).First(&user)
	return c.JSON(user)

}

//Logout function reset the the jwt in cookie to invalidate user credentials
func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:  "jwt",
		Value: "",
		//make the cookie already expired one hour ago
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"msg": "user logout",
	})
}

//UserEmailExists checks if the user emails receives as a paramater in request exists as a username
func UserEmailExists(email string) (models.User, bool) {

	var user models.User
	database.DB.Where("email=?", email).First(&user)
	if user.Id == 0 {
		return user, false
	}
	return user, true

}

//GeneratesJWT receives a models.user object and create the JWT for user auth
func GeneratesJWT(t models.User) (string, error) {
	//the jwt token is an slice of bytes
	key := []byte("ReactGoAPI")
	//claims (privileges) section of the token to add in paylod
	payload := jwt.StandardClaims{
		Subject:   strconv.Itoa(int(t.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}
	//header part of token, encrypton algo
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	//sign the token with key slice of bytes
	tokenStr, err := token.SignedString(key)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil

}
