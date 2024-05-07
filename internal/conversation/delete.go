package conversation

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
	rawId := c.Params("id")
	if rawId == "" {
		return hResp.BadRequestResponse(c, "inform id")
	}

	id, err := primitive.ObjectIDFromHex(rawId)
	if err != nil {
		return hResp.BadRequestResponse(c, err.Error())
	}

	collection := database.Get().Database.Collection("conversations")
	filter := bson.M{
		"_id":        id,
		"deleted_at": nil,
	}
	update := bson.M{"$set": bson.M{"deleted_at": time.Now()}}

	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return hResp.InternalServerErrorResponse(c, err.Error())
	}

	return hResp.SuccessResponse(c, "Conversation deleted successfully")
}
