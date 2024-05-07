package user

import (
	"context"
	"strconv"

	"github.com/flambra/chat/database"
	"github.com/flambra/chat/internal/domain"
	"github.com/flambra/helpers/hResp"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func Read(c *fiber.Ctx) error {
	rawId := c.Params("id")
	if rawId == "" {
		return hResp.BadRequestResponse(c, "inform id")
	}

	id, err := strconv.Atoi(rawId)
	if err != nil {
		return hResp.BadRequestResponse(c, err.Error())
	}

	var user domain.User

	collection := database.Get().Database.Collection("users")
	filter := bson.M{
		"_id":        id,
		"deleted_at": nil,
	}

	err = collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return hResp.NotFoundResponse(c, &user, "User not found")
	}

	return hResp.SuccessResponse(c, &user)
}
