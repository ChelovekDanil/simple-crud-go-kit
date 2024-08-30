package crud

import (
	"context"
	"encoding/json"

	"github.com/chelovekdanil/crud/database"
	"github.com/google/uuid"
)

type Service interface {
	Get(ctx context.Context, id string) (*User, error)
	GetAll(ctx context.Context) ([]User, error)
	Create(ctx context.Context, user User) (string, error)
	Update(ctx context.Context, user User) error
	Delete(ctx context.Context, id string) error
}

type User struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// NewService - create type Service
func NewService() Service {
	return &User{}
}

// Get - returns user by id
func (u *User) Get(ctx context.Context, id string) (*User, error) {
	rdb, err := database.Connect(ctx)
	if err != nil {
		return nil, err
	}
	defer rdb.Close()

	res, err := rdb.Get(ctx, id).Result()
	if err != nil {
		return nil, err
	}

	var data User
	err = json.Unmarshal([]byte(res), &data)
	if err != nil {
		return nil, err
	}
	return &User{Id: id, FirstName: data.FirstName, LastName: data.LastName}, nil
}

// GetAll - return all user
func (u *User) GetAll(ctx context.Context) ([]User, error) {
	rdb, err := database.Connect(ctx)
	if err != nil {
		return nil, err
	}
	defer rdb.Close()

	var cursor uint64
	var keys []string
	for {
		newKey, cursor, err := rdb.Scan(ctx, cursor, "*", 0).Result()
		if err != nil {
			return nil, err
		}
		keys = append(keys, newKey...)

		if cursor == 0 {
			break
		}
	}

	var users []User
	for _, key := range keys {
		res, err := rdb.Get(ctx, key).Result()
		if err != nil {
			return nil, err
		}

		var data User
		err = json.Unmarshal([]byte(res), &data)
		if err != nil {
			return nil, err
		}
		users = append(users, User{Id: key, FirstName: data.FirstName, LastName: data.LastName})
	}
	return users, nil
}

// Create - created a new user
func (u *User) Create(ctx context.Context, user User) (id string, err error) {
	rdb, err := database.Connect(ctx)
	if err != nil {
		return "", err
	}
	defer rdb.Close()

	insertData := User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
	jsonData, err := json.Marshal(insertData)
	if err != nil {
		return "", err
	}
	id = uuid.New().String()
	err = rdb.Set(ctx, id, jsonData, 0).Err()
	if err != nil {
		return "", err
	}
	return id, nil
}

// Update - updates the user
func (u *User) Update(ctx context.Context, user User) error {
	rdb, err := database.Connect(ctx)
	if err != nil {
		return err
	}
	defer rdb.Close()

	insertData := User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
	jsonData, err := json.Marshal(insertData)
	if err != nil {
		return err
	}
	if err = rdb.Set(ctx, user.Id, jsonData, 0).Err(); err != nil {
		return err
	}
	return nil
}

// Delete - delete the user
func (u *User) Delete(ctx context.Context, id string) error {
	rdb, err := database.Connect(ctx)
	if err != nil {
		return err
	}
	defer rdb.Close()

	if err = rdb.Del(ctx, id).Err(); err != nil {
		return err
	}
	return nil
}
