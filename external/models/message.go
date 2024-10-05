package models

type Headers struct {
	RoutingKey string
}

type Message struct {
	body interface{}
	Headers
}
