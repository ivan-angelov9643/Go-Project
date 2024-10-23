package structs

import (
	"github.com/google/uuid"
	"time"
)

type Item struct {
	ID           uuid.UUID
	ListID       uuid.UUID
	Tittle       string
	Description  string
	Tags         []Tag
	Completed    bool
	CreationTime time.Time
}
