package characters

import "github.com/gofiber/fiber/v2"

func GetCharactersController(ctx *fiber.Ctx) error {

	return ctx.SendStatus(200)
}
