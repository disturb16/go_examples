// msgbroker provides a mock implementation of a message broker
package msgbroker

import "log"

type HandlerFunc func(msg string) error

type Processor interface {
	Register(b Broker)
}

type Broker interface {
	Start() error
	Stop() error
	AddProcessor(topic string, f HandlerFunc)
}

type broker struct {
}

func (b *broker) Start() error {
	log.Println("Listener started")
	return nil
}

func (b *broker) Stop() error {
	log.Println("Listener stopped")
	return nil
}

func (b *broker) AddProcessor(topic string, f HandlerFunc) {
	log.Printf("Processor for %s added\n", topic)
}

func New() Broker {
	return &broker{}
}
