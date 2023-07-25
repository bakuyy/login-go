package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.TODO()
	opts := options.Client().ApplyURI("mongodb://localhost:27017/db")

	client, err := mongo.Connect(ctx,opts)
	if err != nil {
		panic(err)
	}

	
	db := client.Database("Info")
	users := db.Collection("Users")
	fmt.Println("users", users)

	defer client.Disconnect(ctx)
}