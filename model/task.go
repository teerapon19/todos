package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID           primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Title        string             `json:"title"`
	Description  string             `json:"description"`
	IsAccomplish bool               `json:"isAccomplish,omitempty"`
	Timestamp    time.Time          `json:"timestamp,omitempty"`
}
