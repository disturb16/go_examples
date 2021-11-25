package sqlserver

import "fmt"

type SqlServer struct{}

func New() *SqlServer {
	return &SqlServer{}
}

func (ss *SqlServer) Find(id int) string {
	return "data from sqlserver"
}

func (ss *SqlServer) Save(data string) error {
	fmt.Println("save data to sqlserver")
	return nil
}
