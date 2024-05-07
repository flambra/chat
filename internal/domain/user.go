package domain

import "time"

type User struct {
	ID        uint       `bson:"_id"`
	Username  string     `bson:"username"`
	DeletedAt *time.Time `bson:"deleted_at,omitempty"`
}
