package domain

import (
	"time"

	"gorm.io/gorm"
)

type Conversation struct {
	ConversationID    uint `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	User      string
	Data      string
}
