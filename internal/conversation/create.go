package conversation

import (
	"context"

	"github.com/flambra/chat/database"
	"github.com/flambra/chat/internal/domain"
	"github.com/flambra/helpers/hResp"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ConversationCreateRequest struct {
	Participants []uint `json:"participants"`
}

func Create(c *fiber.Ctx) error {
	var conversation domain.Conversation
	var request ConversationCreateRequest

	if err := c.BodyParser(&request); err != nil {
		return hResp.BadRequestResponse(c, err.Error())
	}

	participants := make([]domain.User, 2) // Limit 2 user per conversation

	for i, userID := range request.Participants {
		collection := database.Get().Database.Collection("users")
		var user domain.User
		err := collection.FindOne(context.TODO(), bson.M{"_id": userID}).Decode(&user)
		if err != nil {
			return hResp.NotFoundResponse(c, &user, "Not Found")
		}
		participants[i] = user
	}

	conversation = domain.Conversation{
		Participants: participants,
	}

	collection := database.Get().Database.Collection("conversations")
	result, err := collection.InsertOne(context.TODO(), conversation)
	if err != nil {
		return hResp.InternalServerErrorResponse(c, err.Error())
	}

	conversation.ID = result.InsertedID.(primitive.ObjectID)

	return hResp.SuccessCreated(c, &conversation)
	// return hResp.SuccessCreated(c, &result)
}
