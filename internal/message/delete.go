package message

import (
	"context"
	"time"

	"github.com/flambra/chat/database"
	"github.com/flambra/helpers/hResp"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Delete(c *fiber.Ctx) error {
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

	collection := database.Get().Database.Collection("conversations")
	filter := bson.M{
		"_id": convID,
		"messages.message_id": msgID,
	}
	update := bson.M{
		"$set": bson.M{"messages.$.deleted_at": time.Now()},
	}

	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return hResp.InternalServerErrorResponse(c, err.Error())
	}

	return hResp.SuccessResponse(c, "Message deleted successfully")
}