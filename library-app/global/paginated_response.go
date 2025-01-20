package global

type PaginatedResponse[T any] struct {
	PageSize int   `json:"page_size,omitempty"`
	Page     int   `json:"page,omitempty"`
	Count    int64 `json:"count"`
	Data     []T   `json:"data"`
}
