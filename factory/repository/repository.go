package repository

import (
	"fmt"

	"github.com/user/factory-pattern/configuration"
	"github.com/user/factory-pattern/repository/mysql"
	"github.com/user/factory-pattern/repository/sqlserver"
)

type Repository interface {
	Find(id int) string
	Save(data string) error
}

func New(config *configuration.Configuration) (Repository, error) {

	var repo Repository
	var err error

	switch config.Engine {
	case "mysql":
		repo = mysql.New()

	case "sqlserver":
		repo = sqlserver.New()

	default:
		err = fmt.Errorf("invalid engine %s", config.Engine)
	}

	return repo, err
}
