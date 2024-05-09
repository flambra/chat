package watcher

import (
	"context"
	"fmt"
	"log"

	"github.com/flambra/chat/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Init() {
	collection := database.Get().Database.Collection("conversations")

	pipeline := mongo.Pipeline{
		bson.D{{Key: "$match", Value: bson.D{{Key: "operationType", Value: "update"}}}}}
	changeStream, err := collection.Watch(context.TODO(), pipeline)
	if err != nil {
		log.Fatal(err)
	}
	defer changeStream.Close(context.TODO())

	for changeStream.Next(context.TODO()) {
		var change bson.M
		if err := changeStream.Decode(&change); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Change detected:", change)
	}

	if err := changeStream.Err(); err != nil {
		log.Fatal(err)
	}
}
