package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Conversation struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Participants []User             `bson:"participants"`
	Messages     []Message          `bson:"messages,omitempty"`
	DeletedAt    *time.Time         `bson:"deleted_at,omitempty"`
}
