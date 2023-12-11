package database

import (
	"fmt"
	"log"

	"github.com/jeeo/pack-management/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgresDB(config config.DatabaseConfig) *sqlx.DB {
	connectionString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", config.DBHost, config.DBPort, config.DBName, config.DBUser, config.DBPwd)
	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		log.Fatal("failed to connect to database", err)
	}

	return db
}
