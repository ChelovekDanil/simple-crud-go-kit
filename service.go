package crud

import (
	"context"
	"fmt"

	"github.com/chelovekdanil/crud/database"
	_ "github.com/lib/pq"
)

type Service interface {
	Get(ctx context.Context, id string) (*User, error)
	// GetAll(ctx context.Context) ([]User, error)
	// Create(ctx context.Context, user User) error
	// Update(ctx context.Context, user User) error
	// Delete(ctx context.Context, id string) error
}

type User struct {
	Id        string `json:"id"`
	FirstName string `json:"first-name"`
	LastName  string `json:"last-name"`
}

func NewService() Service {
	return &User{}
}

func (s *User) Get(ctx context.Context, id string) (*User, error) {
	conn, err := database.Connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var firstName string
	var lastName string

	row := conn.QueryRowContext(ctx, fmt.Sprintf("SELECT first_name, last_name FROM users WHERE id = %s;", id))
	err = row.Scan(&firstName, &lastName)
	if err != nil {
		return nil, err
	}
	return &User{id, firstName, lastName}, nil
}
