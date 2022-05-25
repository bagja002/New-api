package main

import (
	"log"
	"trashgo/Auth"
	"trashgo/Database"
	"trashgo/Home"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)
func main() {

	Database.Connect()
	app := fiber.New()
	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	app.Post("/auth/register", Auth.Register)
	app.Post("/auth/login", Auth.Login)
	app.Get("/auth/user", Auth.Users)
	app.Get("/auth/logout", Auth.Logout)
	app.Post("/home/laporan", Home.Laporan)
	log.Fatal(app.Listen(":1234"))

}
