package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Database is an implementation of Database interface for mongodb
type Database struct {
	// Context for mongo DB
	Context context.Context
	// Name is name of the database
	Name string

	// Mongo DB client
	client *mongo.Client
}

var DB *Database

func CreateDB(name, uri string) error {

	DB = &Database{Name: name}
	return DB.Connect(uri)
}

func (db *Database) Connect(uri string) error {

	var err error
	// Connecting to mongodb server
	db.Context, _ = context.WithCancel(context.Background())
	db.client, err = mongo.Connect(db.Context, options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	// Checking connection
	err = db.client.Ping(db.Context, readpref.Primary())
	return err
}

func (db *Database) Disconnect() error {

	return db.client.Disconnect(db.Context)
}

func (db *Database) Reset() error {

	return db.client.Database(db.Name).Drop(db.Context)
}

func (db *Database) Collection(name string) (*mongo.Collection, error) {

	return db.client.Database(db.Name).Collection(name), nil
}
