package message

import (
	"context"
	"time"

	"github.com/flambra/chat/database"
	"github.com/flambra/chat/internal/domain"
	"github.com/flambra/helpers/hResp"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageCreateRequest struct {
	ConversationID string      `json:"conversation_id"`
	Path           domain.Path `json:"path"`
	Content        string      `json:"content"`
}

func Create(c *fiber.Ctx) error {
	var message domain.Message
	var request MessageCreateRequest

	if err := c.BodyParser(&request); err != nil {
		return hResp.BadRequestResponse(c, err.Error())
	}

	conversationID, err := primitive.ObjectIDFromHex(request.ConversationID)
	if err != nil {
		return hResp.BadRequestResponse(c, err.Error())
	}

	message = domain.Message{
		MessageID: primitive.NewObjectID(),
		Path:      request.Path,
		Content:   request.Content,
		SentAt:    time.Now(),
	}

	collection := database.Get().Database.Collection("conversations")
	filter := bson.M{"_id": conversationID}
	update := bson.M{"$push": bson.M{"messages": message}}

	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return hResp.InternalServerErrorResponse(c, err.Error())
	}

	return hResp.SuccessCreated(c, nil)
}
