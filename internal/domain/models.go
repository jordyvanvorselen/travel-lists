package domain

import "github.com/google/uuid"

type List struct {
	Id   int
	UUID uuid.UUID
	Name string `form:"name"`
}

type ListItem struct {
	Id       int
	ListId   int
	Location string `form:"location"`
}
