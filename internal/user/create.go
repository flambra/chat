package user

import (
	"context"
	"time"

	"github.com/flambra/chat/database"
	"github.com/flambra/chat/internal/domain"
	"github.com/flambra/helpers/hResp"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserCreateRequest struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
}

func Create(c *fiber.Ctx) error {
	var user domain.User
	var request UserCreateRequest

	if err := c.BodyParser(&request); err != nil {
		return hResp.BadRequestResponse(c, err.Error())
	}

	user = domain.User{
		UserID:    request.UserID,
		Username:  request.Username,
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
	}

	collection := database.Get().Database.Collection("users")
	_, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return hResp.InternalServerErrorResponse(c, err.Error())
	}

	return hResp.SuccessCreated(c, &user)
}
