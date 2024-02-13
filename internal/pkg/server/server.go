package server

import "github.com/gofiber/fiber/v2"

func CreateServer(appName string) *fiber.App {
	app := fiber.New(fiber.Config{
		EnablePrintRoutes: true,
		AppName:           appName,
	})

	return app
}
