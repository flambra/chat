package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	MessageID primitive.ObjectID `bson:"message_id"`
	Path      Path               `bson:"path"`
	Content   string             `bson:"content"`
	SentAt    time.Time          `bson:"sent_at"`
	DeletedAt *time.Time         `bson:"deleted_at,omitempty"`
}

type Path struct {
	From uint `bson:"from"` // UserID of the sending user
	To   uint `bson:"to"`   // UserID of the receiving user
}
