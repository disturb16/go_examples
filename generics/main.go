package main

import (
	"fmt"
	"generics/redisclient"
)

type Numeric interface {
	int | float64 | float32 | int32 | int64
}

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Name     string
	Position string
}

func sum[K string, T Numeric](key K, a T, b T) T {
	fmt.Println(key)
	return a + b
}

func main() {

	// Basic example
	fmt.Println(sum("sum int", 1, 2))
	fmt.Println(sum("sum float", 1.0, 2.0))

	// Redis example
	pp, err := redisclient.Read[[]Person]("persons")
	if err != nil {
		panic(err)
	}

	ee, err := redisclient.Read[[]Employee]("employees")
	if err != nil {
		panic(err)
	}

	fmt.Println("Persons:")
	for _, p := range pp {
		fmt.Printf("%s is %d years old\n", p.Name, p.Age)
	}

	fmt.Println("Employees:")
	for _, e := range ee {
		fmt.Printf("%s is a %s\n", e.Name, e.Position)
	}

}
