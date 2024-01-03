package controller

import (
	"context"
	"fmt"

	"github.com/SatyendraDhamgaye/mongoDbApi/helpers"
	"github.com/SatyendraDhamgaye/mongoDbApi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	//background works in background anutomatically on repeat
	//TODO works in foreground
	client, err := mongo.Connect(context.TODO(), clientOption)

	helpers.ErrorCatcher(err)

	fmt.Println("MongoDB connection success")

	collection = client.Database(dbname).Collection(colName)

	//collection instance
	fmt.Println("Collection instance is ready")
}

//mongodb helper methods should go in seperate file

func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	helpers.ErrorCatcher(err)
	fmt.Println("Inserted 1 movie in the database with Id: ", inserted.InsertedID)
}

func updateOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	helpers.ErrorCatcher(err)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watcher": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	helpers.ErrorCatcher(err)

	fmt.Println("Modified count: ", result.ModifiedCount)
}

func deleteOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	helpers.ErrorCatcher(err)
	filter := bson.M{"_id": id}
	deletecount, err := collection.DeleteOne(context.Background(), filter)
	helpers.ErrorCatcher(err)
	fmt.Println("Movie got deleted with delete count: ", deletecount)
}

func deleteAllMovie(movieId string) {
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}})
	helpers.ErrorCatcher(err)
	fmt.Println("Number of movies deleted: ", deleteResult.DeletedCount)

}
