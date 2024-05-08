package main

import (
	"github.com/Rabbit937/fiber-demo/database"
	"github.com/Rabbit937/fiber-demo/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	database.ConnectDb()

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})

	setupRoutes(app)

	app.Static("/", "./public")

	app.Use(handlers.NotFound)

	app.Listen(":3000")
}
