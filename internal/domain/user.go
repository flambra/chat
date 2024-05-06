package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	UserID    uint               `bson:"_id,omitempty"`
	Username  string             `bson:"username"`
	CreatedAt primitive.DateTime `bson:"created_at"`
}
