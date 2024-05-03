package chat

import (
	"github.com/flambra/chat/internal/client"
	"github.com/flambra/helpers/hResp"
	"github.com/gofiber/fiber/v2"
)

type user struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// RegisterNewUser recebe detalhes de um novo usuário e dispara um evento
func NewUser(c *fiber.Ctx) error {
	newUser := new(user)
	if err := c.BodyParser(newUser); err != nil {
		return hResp.BadRequestResponse(c, err.Error())
	}

	client := client.NewPusher()
	client.Trigger("update", "new-user", newUser)

	return hResp.SuccessResponse(c, &newUser)
}

// PusherAuth autoriza usuários a subscreverem canais privados
func PusherAuth(c *fiber.Ctx) error {
	params := c.Body()

	client := client.NewPusher()
	response, err := client.AuthorizePrivateChannel(params)
	if err != nil {
		return hResp.InternalServerErrorResponse(c, err.Error())
	}

	return hResp.SuccessResponse(c, &response)
}