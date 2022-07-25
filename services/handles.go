package services

import (
	"github.com/aliasaddik/todo-project/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Handles interface {
	CreateTask(*models.Task) error
	GetTask() ([]*models.Task, error)
	EditTask(*models.Task) error
	DeleteTask(*primitive.ObjectID) error
}
