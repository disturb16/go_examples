package main

import (
	"fmt"

	"github.com/disturb16/go-interfaces/users"
)

func main() {
	var frank users.Cashier = users.NewEmployee("Frank")
	var robert users.Admin = users.NewEmployee("Robert")

	total := frank.CalcTotal(90, 65, 93.6)
	fmt.Println(total)

	robert.DeactivateEmployee(frank)

	fmt.Println(frank.CalcTotal(90, 65, 93.6))
}
