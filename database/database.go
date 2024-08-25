package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	// "github.com/chelovekdanil/crud/internal/config"
)

func Connect() (*sql.DB, error) {
	// _ = config.MustLoad().Database

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, fmt.Errorf("cannot connect to db: %s", err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, first_name VARCHAR(50), last_name VARCHAR(50));")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("INSERT INTO users(first_name, last_name) VALUES('danil', 'antonov');")
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
