package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
		cancel()
	}

	collection := client.Database("TaskMaster").Collection("Tasks")

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)

	cur, err := collection.Find(ctx, bson.D{})

	if err != nil {
		log.Fatal(err)
		cancel()
	}

	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v\n\n", result)
	}
}
