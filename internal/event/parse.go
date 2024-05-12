package event

import (
	"time"

	"github.com/flambra/chat/internal/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Parse(change primitive.M) (*domain.Event, error) {
	var event domain.Event

	dk := change["documentKey"].(primitive.M)
	id := dk["_id"].(primitive.ObjectID)
	event.ConversationID = id.Hex()

	ud := change["updateDescription"].(primitive.M)
	uf := ud["updatedFields"].(primitive.M)

	for _, val := range uf {
		message := val.(primitive.M)
		event.Message.Content = message["content"].(string)
		event.Message.MessageID = message["message_id"].(primitive.ObjectID)
		pathData := message["path"].(primitive.M)
		event.Message.Path.From = uint(pathData["from"].(int64))
		event.Message.Path.To = uint(pathData["to"].(int64))		
		sentAt := message["sent_at"].(primitive.DateTime)
		sec := int64(sentAt) / 1000
		nsec := (int64(sentAt) % 1000) * 1000000
		event.Message.SentAt = time.Unix(sec, nsec)
		break
	}

	return &event, nil
}
