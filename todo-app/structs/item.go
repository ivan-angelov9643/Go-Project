package structs

import "time"

type Item struct {
	ID           int
	ListID       int
	Tittle       string
	Description  string
	Tags         []Tag
	Completed    bool
	CreationTime time.Time
}
