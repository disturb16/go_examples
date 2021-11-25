package dbconnection

import (
	"database/sql"
	"fmt"

	"github.com/disturb16/go_example/gorillamux/settings"
	_ "github.com/go-sql-driver/mysql"
)

func New() *sql.DB {
	user := settings.GetEnv("DB_USER", "")
	pass := settings.GetEnv("DB_PASS", "")
	host := settings.GetEnv("DB_HOST", "")
	port := settings.GetEnv("DB_PORT", "")
	dbname := settings.GetEnv("DB_NAME", "")

	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, dbname)

	db, err := sql.Open("mysql", connString)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
