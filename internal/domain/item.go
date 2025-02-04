package domain

import "github.com/google/uuid"

type ItemEntity struct {
	Id          uuid.UUID
	Title       string
	Description string
	Price       float64
	Stock       int
}

type ItemImages struct {
	Id     uuid.UUID
	ItemId uuid.UUID
	Url    string
}
