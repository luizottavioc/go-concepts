package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var uriDb string = "mongodb://localhost:27017"
var db string = "go-concepts"
var collection string = "hello-world"

func connect() (*mongo.Collection, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 10)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uriDb))
	if err != nil { return nil, err }

	collection := client.Database(db).Collection(collection)

	return collection, nil
}

func main() {
	collection, err := connect()
	if err != nil { panic(err) }

	resD, errD := collection.DeleteMany(context.Background(), bson.M{});
	if errD != nil { panic(err) }

	fmt.Printf("Deleted %v document(s)\n", resD.DeletedCount)

	resI, errI := collection.InsertOne(context.Background(), bson.M{"hello": "world"})
	if errI != nil { panic(err) }

	id := resI.InsertedID
	fmt.Printf("Inserted ID: %v\n", id)
}