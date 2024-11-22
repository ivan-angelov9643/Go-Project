package models

import (
	"github.com/google/uuid"
	"time"
)

type Item struct {
	ID           uuid.UUID
	ListID       uuid.UUID
	Title        string
	Description  string
	Tags         []Tag
	Completed    bool
	CreationTime time.Time
}
