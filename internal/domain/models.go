package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type List struct {
	gorm.Model
	UUID  uuid.UUID
	Name  string `form:"name"`
	Items []ListItem
}

type ListItem struct {
	gorm.Model
	ListID   uint
	Location string `form:"location"`
}
