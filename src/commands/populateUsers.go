package main

import (
	"go-react-api/src/database"
	"go-react-api/src/models"

	"github.com/bxcodec/faker/v3"
)

func main() {
	database.Connect()

	for i := 0; i < 30; i++ {
		ambassador := models.User{
			FirstName:    faker.FirstName(),
			LastName:     faker.LastName(),
			Email:        faker.Email(),
			IsAmbassador: true,
		}

		ambassador.SetAndEncryptPassword("12345")
		database.DB.Create(&ambassador)

	}

}
