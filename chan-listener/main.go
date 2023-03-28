package main

import (
	"chan-listener/messanger"
	"log"
	"math/rand"
	"time"
)

type Person struct {
	ID   int
	Name string
}

type Animal struct {
	ID   int
	Name string
}

func handler1(msg any) error {
	log.Println("Message received from topic [persons]: ", msg)
	return nil
}

func handler2(msg any) error {
	log.Println("Message received from topic [animals]: ", msg)
	return nil
}

func main() {
	m := messanger.New()

	m.AddHandler(handler1, "persons")
	m.AddHandler(handler2, "animals")

	go func() {
		for {
			rand.NewSource(time.Now().UnixNano())
			num := rand.Intn(10)

			if num%2 == 0 {
				m.SendMessage(messanger.Message{
					Topic: "persons",
					Data: Person{
						ID:   num,
						Name: "Jhon",
					},
				})
			} else {
				m.SendMessage(messanger.Message{
					Topic: "animals",
					Data: Animal{
						ID:   num,
						Name: "Dog",
					},
				})
			}
		}
	}()

	go m.Start()

	time.Sleep(3 * time.Second)
}
