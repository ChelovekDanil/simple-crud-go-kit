package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect return connection in database
func Connect(ctx context.Context) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	conn, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
