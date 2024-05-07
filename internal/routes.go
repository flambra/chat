// internal/routes.go

package internal

import (
	"os"

	"github.com/flambra/chat/internal/conversation"
	"github.com/flambra/chat/internal/message"
	"github.com/flambra/chat/internal/user"
	"github.com/gofiber/fiber/v2"
)

func InitializeRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"project":     os.Getenv("PROJECT"),
			"environment": os.Getenv("ENV"),
			"version":     os.Getenv("BUILD_VERSION"),
		})
	})

	app.Post("/user", user.Create)
	app.Get("/user/:id", user.Read)
	app.Put("/user/:id", user.Update)
	app.Delete("/user/:id", user.Delete)

	app.Post("/conversation", conversation.Create)
	app.Get("/conversation/:id", conversation.Read)
	// app.Put("/conversation/:id", conversation.Update)
	app.Delete("/conversation/:id", conversation.Delete)

	app.Post("/message", message.Create)
	app.Get("/message/:id", message.Read)
	app.Put("/message/:id", message.Update)
	app.Delete("/message/:id", message.Delete)
}
