package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Worries struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	Title      string             `json:"title"`
	Content    string             `json:"content"`
	Nickname   string             `json:"nickname"`
	AiAdvice   string             `json:"ai_advice"`
	CategoryId primitive.ObjectID `json:"category_id" bson:"category_id"`
	CreatedAt  time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt  time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
