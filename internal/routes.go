// internal/routes.go

package internal

import (
	"github.com/flambra/chat/internal/server"
	"github.com/flambra/chat/internal/user"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func InitializeRoutes(app *fiber.App) {
	app.Post("/user", user.Create)

	app.Use("/ws", server.Upgrade)

	// Rota WebSocket efetiva, manipulando o upgrade
	app.Get("/ws", websocket.New(server.WebSocketHandler()))
}
