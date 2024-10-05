package models

import (
	"errors"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type MessageBus struct {
	connection *amqp.Connection
	URL        string
	UserName   string
	Password   string
	Port       string
	queues     map[string]*Queue
	exchanges  map[string]*Exchange
}

func NewMessageBus(username, password, port, brokerIP string) *MessageBus {
	if username == "" {
		username = "guest"
	}
	if password == "" {
		password = "guest"
	}
	if port == "" {
		port = "5672"
	}
	if brokerIP == "" {
		brokerIP = "localhost"
	}
	url := fmt.Sprintf("amqp://%s:%s@%s:%s/", username, password, brokerIP, port)
	return &MessageBus{
		URL:      url,
		UserName: username,
		Password: password,
		Port:     port,
	}
}

func (mb *MessageBus) ConnectToMessageBus() error {
	if mb == nil {
		return errors.New("message bus object missing")
	}
	var err error
	mb.connection, err = amqp.Dial(mb.URL)
	if err != nil {
		return err
	}
	return nil
}
