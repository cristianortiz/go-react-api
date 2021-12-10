package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {

	//_, err := gorm.Open(mysql.Open("root:invernalia2013@tcp(127.0.0.1:3306)/ambassador"), &gorm.Config{})
	_, err := gorm.Open(mysql.Open("root:root@tcp(db:3306)/ambassador"), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the Database!!")
	}

}
