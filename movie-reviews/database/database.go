package database

import "github.com/jmoiron/sqlx"

type Config struct {
	Host   string `json:"database.host"`
	Port   string `json:"database.port"`
	User   string `json:"database.user"`
	Pass   string `json:"database.password"`
	DbName string `json:"database.dbname"`
}

func New() *sqlx.DB {}
