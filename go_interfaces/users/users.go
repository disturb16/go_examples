package users

import (
	"fmt"
	"math/rand"
)

type User struct {
	Id   int
	Name string
}

type Employee struct {
	User
	Active bool
}

type Cashier interface {
	CalcTotal(item ...float32) float32
	deactivate()
	reactivate()
}

type Admin interface {
	DeactivateEmployee(c Cashier)
}

func NewEmployee(name string) *Employee {
	return &Employee{
		User: User{
			Id:   rand.Intn(1000),
			Name: name,
		},
		Active: true,
	}
}

func (e *Employee) CalcTotal(item ...float32) float32 {

	if !e.Active {
		fmt.Println("Cashier deactivated")
		return 0
	}

	var sum float32

	for _, i := range item {
		sum += i
	}

	return sum * 1.15
}

func (e *Employee) deactivate() {
	e.Active = false
}

func (e *Employee) reactivate() {
	e.Active = true
}

func (e *Employee) DeactivateEmployee(c Cashier) {
	c.deactivate()
}
