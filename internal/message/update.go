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
	rawId := c.Params("id")
	if rawId == "" {
		return hResp.BadRequestResponse(c, "inform id")
	}

	id, err := primitive.ObjectIDFromHex(rawId)
	if err != nil {
		return hResp.BadRequestResponse(c, err.Error())
	}

	var request MessageUpdateRequest

	if err := c.BodyParser(&request); err != nil {
		return hResp.BadRequestResponse(c, err.Error())
	}

	collection := database.Get().Database.Collection("messages")
	filter := bson.M{
		"_id":        id,
		"deleted_at": nil,
	}
	update := bson.M{"$set": bson.M{"content": request.Content}}

	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return hResp.InternalServerErrorResponse(c, err.Error())
	}

	return hResp.SuccessResponse(c, "Message updated successfully")
}
