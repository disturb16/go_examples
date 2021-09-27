package main

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
)

func main() {

	wg := &sync.WaitGroup{}
	IDsChan := make(chan string)
	FakeIDsChan := make(chan string)
	closedChans := make(chan int)

	wg.Add(3)

	go generateIDS(wg, IDsChan, closedChans)
	go generateFakeIDs(wg, FakeIDsChan, closedChans)

	go logIDs(wg, IDsChan, FakeIDsChan, closedChans)

	wg.Wait()
}

func generateFakeIDs(wg *sync.WaitGroup, fakeIDsChan chan<- string, closedChannels chan<- int) {
	for i := 0; i < 100; i++ {
		id := uuid.New()
		fakeIDsChan <- fmt.Sprintf("%d. %s", i+1, id.String())
	}

	close(fakeIDsChan)
	closedChannels <- 1

	wg.Done()
}

func generateIDS(wg *sync.WaitGroup, idsChan chan<- string, closedChannels chan<- int) {

	for i := 0; i < 100; i++ {
		id := uuid.New()
		idsChan <- fmt.Sprintf("%d. %s", i+1, id.String())
	}

	close(idsChan)
	closedChannels <- 1

	wg.Done()
}

func logIDs(wg *sync.WaitGroup, idsChan <-chan string, fakeIDsChan <-chan string, closedChannels chan int) {

	closedCounter := 0

	for {
		select {
		case id, ok := <-idsChan:
			if ok {
				fmt.Println("ID:", id)
			}

		case id, ok := <-fakeIDsChan:
			if ok {
				fmt.Println("FAKE ID:", id)
			}

		case count, ok := <-closedChannels:
			if ok {
				closedCounter += count
			}
		}

		if closedCounter == 2 {
			close(closedChannels)
			break
		}
	}

	wg.Done()
}
