package main

import "fmt"

func main() {
	fruits := []string{
		"Manzana",
		"Banano",
		"Sandia",
		"Melon",
	}

	for _, fruit := range fruits {
		index := len(fruit) - 1
		letter := fruit[index:]

		if letter == "a" {
			continue
		}

		fmt.Println("Fruit:", fruit)
	}
}
