package main

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	//_, err := gorm.Open(mysql.Open("root:invernalia2013@tcp(127.0.0.1:3306)/ambassador"), &gorm.Config{})
	_, err := gorm.Open(mysql.Open("root:root@tcp(db:3306)/ambassador"), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the Database!!")
	}
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello, world")
	})
	app.Listen(":8000")
}
