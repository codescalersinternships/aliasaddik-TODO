package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`

	Done bool `json:"done" bson:"done"`

	Title string `json:"title" bson:"title"`
}
