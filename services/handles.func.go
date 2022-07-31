package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/aliasaddik/todo-project/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type HandlesFunc struct {
	task *mongo.Collection
	ctx  context.Context
}

//acts like a constructor
func NewHandle(task *mongo.Collection, ctx context.Context) Handles {
	return &HandlesFunc{
		task: task,
		ctx:  ctx,
	}
}
func (handle *HandlesFunc) CreateTask(task *models.Task) error {
	_, err := handle.task.InsertOne(handle.ctx, task)
	return err
}

func (handle *HandlesFunc) GetTask() ([]*models.Task, error) {

	var users []*models.Task
	cursor, err := handle.task.Find(handle.ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(handle.ctx) {
		var user models.Task
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(handle.ctx)

	if len(users) == 0 {
		return nil, errors.New("documents not found")
	}
	return users, nil
}
func (handle *HandlesFunc) EditTask(iTask *models.Task) error {
	filter := bson.D{primitive.E{Key: "_id", Value: iTask.ID}}
	fmt.Print(iTask.ID.String())
	update := bson.D{primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "title", Value: iTask.Title}, primitive.E{Key: "_id", Value: iTask.ID}, primitive.E{Key: "done", Value: iTask.Done}}}}
	result, _ := handle.task.UpdateOne(handle.ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("no matched document found for update")
	}
	return nil
}

func (handle *HandlesFunc) DeleteTask(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": objID}
	fmt.Print("the Id i got", objID)
	result, _ := handle.task.DeleteOne(handle.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("no matched document found for delete")
	}
	return nil

}
