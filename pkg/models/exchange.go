package models

import "github.com/ankit081998/rabbitmq-infra/pkg/models"

type ExchangeType string

const (
	DIRECT  ExchangeType = "direct"
	FANOUT  ExchangeType = "fanout"
	HEADER  ExchangeType = "header"
	DEFAULT ExchangeType = ""
	TOPIC   ExchangeType = "TOPIC"
)

type Exchange struct {
	Name          string
	RoutingKey    string
	CreatedAt     uint64
	BoundedQueues map[string]*Queue
	Type          ExchangeType
}

func (exchange *Exchange) GetBoundedQueues(ids ...string) []*Queue {
	queues := make([]*Queue, 0)
	if len(ids) == 0 {
		for _, queueData := range exchange.BoundedQueues {
			queues = append(queues, queueData)
		}
	} else {
		for _, id := range ids {
			queues = append(queues, exchange.BoundedQueues[id])
		}
	}
	return queues
}

func NewExchange(name, routingKey string, Type ExchangeType, queues ...models.Queue) *Exchange {
	newExchange := &Exchange{
		Name:       name,
		RoutingKey: routingKey,
		Type:       Type,
	}
	if len(queues) > 0 {
		for _, queue := range queues {
			newExchange.BoundedQueues[queue.Name] = queue
		}
	}
	
	return newExchange
}
