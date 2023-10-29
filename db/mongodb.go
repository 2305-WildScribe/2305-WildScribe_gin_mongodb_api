package db

import (
    "context"
	"log"
    "gin-mongo-api/configs"
	"go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client = ConnectToMongoDB()
var database = configs.DatabaseEnvironment()

func ConnectToMongoDB() *mongo.Client {
    // Set client options
    clientOptions := options.Client().ApplyURI(configs.EnvMongoURI())
    // Create a MongoDB client
    client, err := mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        log.Printf("Error connecting to MongoDB: %v", err)
    }
    // Ping the server to ensure the connection is established
    err = client.Ping(context.Background(), nil)
    if err != nil {
        log.Printf("Error pinging MongoDB: %v", err)
    }
    // Success!
    log.Println("Connected to MongoDB")
    return client
}

// Get a specfic collection
func GetCollection(collectionName string) *mongo.Collection {
    collection := client.Database(database).Collection(collectionName)
    return collection
}

func SetDataBase(db_string string) {
    database = db_string
}

