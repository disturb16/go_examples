package mysql

import "fmt"

type Mysql struct{}

func New() *Mysql {
	return &Mysql{}
}

func (m *Mysql) Find(id int) string {
	return "data from mysql"
}

func (m *Mysql) Save(data string) error {
	fmt.Println("save data to mysql")
	return nil
}
