package message

import (
	"fmt"

	"github.com/flambra/chat/internal/client"
	"github.com/flambra/helpers/hResp"
	"github.com/gofiber/fiber/v2"
)

type MessageSendRequest struct {
	User string `json:"user"`
	Data string `json:"data"`
}

func Send(c *fiber.Ctx) error {
	// var message domain.Message
	var request MessageSendRequest

	if err := c.BodyParser(&request); err != nil {
		return hResp.BadRequestResponse(c, err.Error())
	}

	pusherClient := client.NewPusher()

	err := pusherClient.Trigger("chat", "message", request.Data)
	if err != nil {
		fmt.Println(err.Error())
		hResp.BadRequestResponse(c, err.Error())
	}

	return hResp.SuccessResponse(c, "message sent")
}
