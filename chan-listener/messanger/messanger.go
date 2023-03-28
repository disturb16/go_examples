package messanger

import (
	"log"
	"sync"
)

type HandlerFunc func(any) error

type Message struct {
	Topic string
	Data  any
}

type Messanger struct {
	mainChan  chan Message
	handlers  map[string]HandlerFunc
	buffer    []Message
	wg        *sync.WaitGroup
	isRunning bool
}

func New() *Messanger {
	return &Messanger{
		mainChan: make(chan Message),
		handlers: map[string]HandlerFunc{},
	}
}

func (m *Messanger) AddHandler(h HandlerFunc, topic string) {
	m.handlers[topic] = h
}

func (m *Messanger) SendMessage(msg Message) {
	if !m.isRunning {
		m.buffer = append(m.buffer, msg)
		return
	}

	m.mainChan <- msg
}

func (m *Messanger) Listen() {
	for msg := range m.mainChan {
		h := m.handlers[msg.Topic]
		err := h(msg)
		if err != nil {
			log.Panic(err)
		}
	}
}

func (m *Messanger) Start() {
	m.wg = &sync.WaitGroup{}
	m.wg.Add(1)
	go m.Listen()

	m.isRunning = true

	for _, msg := range m.buffer {
		m.mainChan <- msg
	}

	m.wg.Wait()
}

func (m *Messanger) Stop() {
	close(m.mainChan)
	m.isRunning = false
	m.wg.Done()
}
