package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/flambra/chat/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var instance *domain.Mongo

func New() error {
	uri := os.Getenv("MONGO_URI")
	db := os.Getenv("MONGO_DB")

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("Error pinging MongoDB: %v", err)
	}
	log.Println("Connected to MongoDB!")

	instance = &domain.Mongo{
		Client:   client,
		Database: client.Database(db),
	}

	return nil
}

func Get(collection string) *mongo.Collection {
	if instance == nil {
		max_attempts := 3
		for attempt := 0; attempt < max_attempts; attempt++ {
			log.Println("retrying connect... attempt: ", attempt)
			New()
			if instance != nil {
				return instance.Database.Collection(collection)
			}
		}
	}
	return instance.Database.Collection(collection)
}

func Disconnect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := instance.Client.Disconnect(ctx); err != nil {
		return fmt.Errorf("error disconnecting from MongoDB: %v", err)
	}

	log.Println("Disconnected from MongoDB")
	return nil
}
