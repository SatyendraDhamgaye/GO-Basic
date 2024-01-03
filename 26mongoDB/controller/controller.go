package controller

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/mongocrypt/options"
)

// connection string from created database from mongodb
const connectionString = "mongodb+srv://satye:Lavra@golangprac.3cu4y7i.mongodb.net/?retryWrites=true&w=majority"

// database details
const dbname = "netflix"
const colName = "watchlist"

// importent part
var collection *mongo.Collection

// connection with mongo DB
func init() {
	//client option list like jdbs in java
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongoDb
	mongo.Connect(context.TODO(), clientOption)
}
