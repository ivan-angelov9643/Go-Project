package models

import (
	"github.com/google/uuid"
	"time"
)

type List struct {
	ID           uuid.UUID
	Name         string
	Description  string
	CreationTime time.Time
	Items        []Item
}
