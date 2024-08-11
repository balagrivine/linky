package config

import (
	"database/sql"
	"log"
	"os"

	"github.com/balagrivine/linky/internal/database"

	_"github.com/lib/pq"
)

type DBConfig struct {
	DB *database.Queries
}

// Function to initialize a database connection
func InitConfig() (*DBConfig ,error) {

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == ""{
		log.Println("DB_URL not set as an environmental variable")
	}

	// Make a database connection
	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		return nil, err
	}

	if err = Ping(conn); err != nil {
		return nil, err
	}

	queries := database.New(conn)

	return &DBConfig{
		DB: queries,
	}, nil
}
