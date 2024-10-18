package domain

import "github.com/google/uuid"

type List struct {
	Id   uuid.UUID
	Name string `form:"name"`
}
