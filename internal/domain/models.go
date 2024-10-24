package domain

import "github.com/google/uuid"

type List struct {
	Id   int
	UUID uuid.UUID
	Name string `form:"name"`
}
