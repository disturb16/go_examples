package database

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

// DBConfig represents the database configuration.
type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

// LoadConfig loads the database configuration from the environment.
//
// For simplicity, we are using hard-coded values.
func LoadConfig() *DBConfig {

	envPort := os.Getenv("DB_PORT")
	port, err := strconv.Atoi(envPort)
	if err != nil {
		panic(err)
	}

	return &DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     port,
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}
}

// NewPostgresDB creates a new Postgres database connection.
func NewPostgresDB(ctx context.Context, config *DBConfig) *sqlx.DB {
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.DBName,
	)

	db, err := sqlx.ConnectContext(ctx, "postgres", connString)
	if err != nil {
		panic(err)
	}

	sqlx.BindDriver("postgres", sqlx.DOLLAR)
	return db
}
