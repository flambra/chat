package message

import (
	"context"

	"github.com/flambra/chat/database"
	"github.com/flambra/chat/internal/domain"
	"github.com/flambra/helpers/hResp"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Read(c *fiber.Ctx) error {
	conversationID := c.Params("id")
	messageID := c.Params("message_id")

	if conversationID == "" || messageID == "" {
		return hResp.BadRequestResponse(c, "Both conversation ID and message ID must be informed")
	}

	convID, err := primitive.ObjectIDFromHex(conversationID)
	if err != nil {
		return hResp.BadRequestResponse(c, "Invalid conversation ID format")
	}

	msgID, err := primitive.ObjectIDFromHex(messageID)
	if err != nil {
		return hResp.BadRequestResponse(c, "Invalid message ID format")
	}

	var message domain.Message

	collection := database.Get("messages")
	filter := bson.M{
		"_id":                 convID,
		"messages.message_id": msgID,
		"deleted_at":          nil,
	}

	err = collection.FindOne(context.TODO(), filter).Decode(&message)
	if err != nil {
		return hResp.NotFoundResponse(c, &message, "Message not found")
	}

	return hResp.SuccessResponse(c, &message)
}
