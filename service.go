package crud

import (
	"context"

	"github.com/chelovekdanil/crud/database"
	_ "github.com/lib/pq"
)

type Service interface {
	Get(ctx context.Context, id string) (*User, error)
	GetAll(ctx context.Context) ([]User, error)
	Create(ctx context.Context, user User) (int, error)
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
	conn, err := database.Connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var firstName string
	var lastName string

	quary := "SELECT first_name, last_name FROM users WHERE id = $1;"

	err = conn.QueryRowContext(ctx, quary, id).Scan(&firstName, &lastName)
	if err != nil {
		return nil, err
	}
	return &User{id, firstName, lastName}, nil
}

// GetAll - return all user
func (u *User) GetAll(ctx context.Context) ([]User, error) {
	conn, err := database.Connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var users []User

	rows, err := conn.QueryContext(ctx, "SELECT id, first_name, last_name FROM users;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id string
		var firstName string
		var lastName string

		rows.Scan(&id, &firstName, &lastName)
		users = append(users, User{id, firstName, lastName})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// Create - created a new user
func (u *User) Create(ctx context.Context, user User) (int, error) {
	conn, err := database.Connect()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	var id int
	query := "INSERT INTO users(first_name, last_name) VALUES($1, $2) RETURNING id"

	err = conn.QueryRowContext(ctx, query, user.FirstName, user.LastName).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// Update - updates the user
func (u *User) Update(ctx context.Context, user User) error {
	conn, err := database.Connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	quary := "UPDATE users SET first_name = $1, last_name = $2 WHERE id = $3;"

	err = conn.QueryRowContext(ctx, quary, user.FirstName, user.LastName, user.Id).Scan()
	if err != nil {
		return err
	}

	return nil
}

// Delete - delete the user
func (u *User) Delete(ctx context.Context, id string) error {
	conn, err := database.Connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	quary := "DELETE FROM users WHERE id = $1;"

	err = conn.QueryRowContext(ctx, quary, id).Scan()
	if err != nil {
		return err
	}

	return nil
}
