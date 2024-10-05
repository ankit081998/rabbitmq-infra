package models

type Queue struct {
	Name string
	Durability bool
	AutoAcknowledgement bool
	CreatedAt uint64
}
