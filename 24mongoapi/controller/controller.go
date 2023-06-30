package controller

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

const connectionString = "mongodb+srv://vaibhavng7:vaibhav16@cluster0.kwoel2e.mongodb.net/?retryWrites=true&w=majority"
const dbname = "netflix"
const collname = "watchlist"

// Most important
var collection *mongo.Collection

//connecting with mongodb

func init() {
	//client option
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connection success")

	collection = client.Database(dbname).Collection(collname)

	//collection insctance
	fmt.Println("Collection instance is ready")
}
