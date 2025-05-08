package controller

import (
	"github.com/Skliar-Il/People-API/internal/container/initializer"
	"github.com/gofiber/fiber/v3"
)

func NewController(app *fiber.App, services *initializer.ServiceList) {
	api := app.Group("/api")

	api.Get("/ping", func(c fiber.Ctx) error {
		return c.Status(200).JSON("pong")
	})

	peopleHandler := NewPeopleHandler(services.PeopleService)
	peopleController := api.Group("/people")
	{
		peopleController.Post("", peopleHandler.CreatePeople)
		peopleController.Get("/:id", peopleHandler.GetPeople)
		peopleController.Put("/:id", peopleHandler.UpdatePeople)
		peopleController.Delete("/:id", peopleHandler.DeletePeople)
	}
}
