package db

import (
	"context"
	"os"
	"time"

	"go-rest-api/pkg/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Client is mongo database
var Client *mongo.Client

// Connect to mongo & create client
func Connect() {
	uri := os.Getenv("MONGO_URI")

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri)
	client, e := mongo.Connect(ctx, clientOptions)

	utils.CheckErr(e)
	Client = client

	e = Client.Ping(ctx, nil)
	utils.CheckErr(e)
	e = Client.Ping(context.TODO(), nil)
	utils.CheckErr(e)
}
