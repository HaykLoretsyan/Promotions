package main

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func CreateRepository() error {

	c, err := DB.Collection("promotions")
	if err != nil {
		return err
	}

	_, err = c.Indexes().CreateOne(
		DB.Context,
		mongo.IndexModel{
			Keys:    bson.M{"id": 1},
			Options: options.Index().SetUnique(true)})
	if err != nil {
		return err
	}

	collection = c
	return nil
}

func Get(id string) (Model, error) {

	result := Model{}
	err := collection.FindOne(DB.Context, bson.D{bson.E{Key: "id", Value: id}}).Decode(&result)
	if err != nil {
		return Model{}, err
	}
	return result, nil
}

func SaveList(models []Model) error {

	_, err := collection.InsertMany(
		DB.Context,
		func() []interface{} {
			result := make([]interface{}, len(models))
			for i, model := range models {
				result[i] = model
			}
			return result
		}())
	return err
}

func RemoveAll() error {

	_, err := collection.DeleteMany(DB.Context, bson.D{})
	return err
}
