package user

import (
	"context"
	"strconv"

	"github.com/flambra/chat/database"
	"github.com/flambra/helpers/hResp"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type UserUpdateRequest struct {
	Username string `json:"username"`
}

func Update(c *fiber.Ctx) error {
	rawId := c.Params("id")
	if rawId == "" {
		return hResp.BadRequestResponse(c, "inform id")
	}

	id, err := strconv.Atoi(rawId)
	if err != nil {
		return hResp.BadRequestResponse(c, err.Error())
	}

	var request UserUpdateRequest

	if err := c.BodyParser(&request); err != nil {
		return hResp.BadRequestResponse(c, err.Error())
	}

	collection := database.Get("users")
	filter := bson.M{
		"_id":        id,
		"deleted_at": nil,
	}
	update := bson.M{"$set": bson.M{"username": request.Username}}

	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return hResp.InternalServerErrorResponse(c, err.Error())
	}

	return hResp.SuccessResponse(c, "User updated successfully")
}
