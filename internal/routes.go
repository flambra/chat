package internal

import (
	"os"

	"github.com/flambra/chat/internal/chat"
	"github.com/flambra/chat/internal/message"
	"github.com/flambra/chat/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func InitializeRoutes(app *fiber.App) {
	// app.Static("/", "./public")

	app.Get("/info", middleware.Auth, func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"project":     os.Getenv("PROJECT"),
			"environment": os.Getenv("ENV"),
			"version":     os.Getenv("BUILD_VERSION"),
		})
	})

	app.Post("/message/send", middleware.Auth, message.Send)

	app.Post("/new/user", chat.NewUser)
	app.Post("/pusher/auth", chat.PusherAuth)

}
