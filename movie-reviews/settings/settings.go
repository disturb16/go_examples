package settings

import (
	"encoding/json"
	"log"

	"github.com/disturb16/go_examples/movie-reviews/database"
	"github.com/joho/godotenv"
)

type Settings struct {
	Port     string `json:"app.port"`
	Database database.Config
}

func New() *Settings {
	config, err := godotenv.Read(".env")
	if err != nil {
		log.Panic(err)
	}

	content, err := json.Marshal(config)
	if err != nil {
		log.Panic(err)
	}

	app := Settings{}

	err = json.Unmarshal(content, &app)
	if err != nil {
		log.Panic(err)
	}

	dbConfig := database.Config{}

	err = json.Unmarshal(content, &dbConfig)
	if err != nil {
		log.Panic(err)
	}

	app.Database = dbConfig

	return &app
}
