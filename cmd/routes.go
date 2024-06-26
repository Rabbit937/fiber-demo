package main

import (
	"github.com/Rabbit937/fiber-demo/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.ListFacts)

	app.Get("/fact", handlers.NewFactView)
	app.Post("/fact", handlers.CreateFact)

	app.Get("/fact/:id", handlers.ShowFact)

	app.Get("/fact/:id/edit", handlers.EditFact)

	app.Patch("/fact/:id", handlers.UpdateFact)

	// Delete fact
	app.Delete("/fact/:id",handlers.DeleteFact)

}
