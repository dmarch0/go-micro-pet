package characters

import "github.com/gofiber/fiber/v2"

func ApplyRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	domain := v1.Group("/characters")

	domain.Get("/", GetCharactersController)
}
