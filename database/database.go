package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	defaultMongoURI = "mongodb://localhost:27017"
	MongoURI        = "MONGO_URI"
)

// Connect return connection in database
func Connect(ctx context.Context) (*mongo.Client, error) {
	var mongoUri string
	if mongoUri = os.Getenv(MongoURI); mongoUri == "" {
		mongoUri = defaultMongoURI
	}
	log.Println(mongoUri)
	clientOptions := options.Client().ApplyURI(mongoUri)

	conn, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
