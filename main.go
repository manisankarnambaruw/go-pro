package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/manisankarnambaruw/go-pro/controllers"
	"github.com/manisankarnambaruw/go-pro/routes"
)

func main() {
	app := fiber.New()

	app.Static("/", "./home.html")

	go controllers.RunHub()

	routes.Setup(app)

	app.Listen(":3000")
}
