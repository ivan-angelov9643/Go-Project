package structs

import "time"

type List struct {
	ID           int
	Name         string
	Description  string
	CreationTime time.Time
	Items        []Item
}
