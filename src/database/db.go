package database

import (
	"go-react-api/src/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect function use gorm functions to connect to DB
func Connect() {
	var err error
	//_, err := gorm.Open(mysql.Open("root:invernalia2013@tcp(127.0.0.1:3306)/ambassador"), &gorm.Config{})
	DB, err = gorm.Open(mysql.Open("root:root@tcp(db:3306)/ambassador"), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the Database!!")
	}

}

//Automigrate function  calls the gorm.automigrate function
func AutoMigrate() {
	DB.AutoMigrate(models.User{})
}
