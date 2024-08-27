package database

import (
	"database/sql"
	"fmt"
	"os"
)

// Connect return connection in database
func Connect() (*sql.DB, error) {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, fmt.Errorf("cannot connect to db: %s", err)
	}

	return db, nil
}
