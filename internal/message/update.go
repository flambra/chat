package message

import (
	"context"

	"github.com/flambra/chat/database"
	"github.com/flambra/helpers/hResp"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageUpdateRequest struct {
	Content string `json:"content"`
}

func Update(c *fiber.Ctx) error {
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

	var request MessageUpdateRequest

	if err := c.BodyParser(&request); err != nil {
		return hResp.BadRequestResponse(c, err.Error())
	}

	collection := database.Get("messages")
	filter := bson.M{
		"_id":                 convID,
		"messages.message_id": msgID,
		"deleted_at":          nil,
	}
	update := bson.M{"$set": bson.M{"content": request.Content}}

	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return hResp.InternalServerErrorResponse(c, err.Error())
	}

	return hResp.SuccessResponse(c, "Message updated successfully")
}
