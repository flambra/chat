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
	rawId := c.Params("id")
	if rawId == "" {
		return hResp.BadRequestResponse(c, "inform id")
	}

	id, err := primitive.ObjectIDFromHex(rawId)
	if err != nil {
		return hResp.BadRequestResponse(c, err.Error())
	}

	var message domain.Message

	collection := database.Get().Database.Collection("messages")
	filter := bson.M{
		"_id":        id,
		"deleted_at": nil,
	}

	err = collection.FindOne(context.TODO(), filter).Decode(&message)
	if err != nil {
		return hResp.NotFoundResponse(c, &message, "Message not found")
	}

	return hResp.SuccessResponse(c, &message)
}
