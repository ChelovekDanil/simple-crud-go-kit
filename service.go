package crud

import (
	"context"

	"github.com/chelovekdanil/crud/database"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type Service interface {
	Get(ctx context.Context, id string) (*User, error)
	GetAll(ctx context.Context) ([]User, error)
	Create(ctx context.Context, user User) (string, error)
	Update(ctx context.Context, user User) error
	Delete(ctx context.Context, id string) error
}

type User struct {
	Id        string `json:"id" bson:"_id"`
	FirstName string `json:"firstName" bson:"firstName"`
	LastName  string `json:"lastName" bson:"lastName"`
}

// NewService - create type Service
func NewService() Service {
	return &User{}
}

const (
	COLLECTION_NAME = "users"
	DBNAME_ENV      = "DBNAME"
)

// Get - returns user by id
func (u *User) Get(ctx context.Context, id string) (*User, error) {
	conn, err := database.Connect(ctx)
	if err != nil {
		return nil, err
	}
	defer func(ctx context.Context) {
		conn.Disconnect(ctx)
	}(ctx)

	collection := conn.Database("crud").Collection(COLLECTION_NAME)
	filter := bson.D{{Key: "_id", Value: id}}
	var user User

	err = collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &User{Id: user.Id, FirstName: user.FirstName, LastName: user.LastName}, nil
}

// GetAll - return all user
func (u *User) GetAll(ctx context.Context) ([]User, error) {
	conn, err := database.Connect(ctx)
	if err != nil {
		return nil, err
	}
	defer func(ctx context.Context) {
		conn.Disconnect(ctx)
	}(ctx)

	collection := conn.Database("crud").Collection(COLLECTION_NAME)
	filter := bson.D{{}}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var users []User
	if err = cursor.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, nil
}

// Create - created a new user
func (u *User) Create(ctx context.Context, user User) (string, error) {
	conn, err := database.Connect(ctx)
	if err != nil {
		return "", err
	}
	defer func(ctx context.Context) {
		conn.Disconnect(ctx)
	}(ctx)

	collection := conn.Database("crud").Collection(COLLECTION_NAME)
	id := uuid.New()

	_, err = collection.InsertOne(ctx, User{Id: id.String(), FirstName: user.FirstName, LastName: user.LastName})
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

// Update - updates the user
func (u *User) Update(ctx context.Context, user User) error {
	conn, err := database.Connect(ctx)
	if err != nil {
		return err
	}
	defer func(ctx context.Context) {
		conn.Disconnect(ctx)
	}(ctx)

	collection := conn.Database("crud").Collection(COLLECTION_NAME)
	filter := bson.M{"_id": user.Id}
	update := bson.M{
		"$set": bson.M{
			"firstName": user.FirstName,
			"lastName":  user.LastName,
		},
	}

	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

// Delete - delete the user
func (u *User) Delete(ctx context.Context, id string) error {
	conn, err := database.Connect(ctx)
	if err != nil {
		return err
	}
	defer func(ctx context.Context) {
		conn.Disconnect(ctx)
	}(ctx)

	collection := conn.Database("crud").Collection(COLLECTION_NAME)
	filter := bson.M{"_id": id}

	_, err = collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
