package mongodb

import (
	"context"
	"time"

	"github.com/teerapon19/todos/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func TaskGetSingle(id primitive.ObjectID) (*model.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var task model.Task
	err := Database.Collection("tasks").FindOne(ctx, bson.M{
		"_id": id,
	}).Decode(&task)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func TaskGetAll() ([]model.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := Database.Collection("tasks").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	tasks := []model.Task{}
	for cursor.Next(ctx) {
		var task model.Task
		if err = cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func TaskInsertNew(task model.Task) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := Database.Collection("tasks").InsertOne(ctx, bson.M{
		"title":        task.Title,
		"description":  task.Description,
		"isAccomplish": false,
		"timestamp":    time.Now(),
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func TaskUpdateEdited(task model.Task) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := Database.Collection("tasks").UpdateOne(ctx, bson.M{
		"_id": task.ID,
	}, bson.M{
		"$set": bson.M{
			"title":       task.Title,
			"description": task.Description,
		},
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func TaskMarkAction(id primitive.ObjectID, isAccomplished bool) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := Database.Collection("tasks").UpdateOne(ctx, bson.M{
		"_id": id,
	}, bson.M{
		"$set": bson.M{
			"isAccomplish": isAccomplished,
		},
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func TaskDelete(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := Database.Collection("tasks").DeleteOne(ctx, bson.M{
		"_id": id,
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}
