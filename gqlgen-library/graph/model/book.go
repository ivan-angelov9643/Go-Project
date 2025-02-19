package model

type Book struct {
	ID              string    `json:"id"`
	Title           string    `json:"title"`
	Year            int32     `json:"year"`
	AuthorID        string    `json:"authorId"`
	CategoryID      string    `json:"categoryId"`
	Category        *Category `json:"category"`
	TotalCopies     int32     `json:"totalCopies"`
	AvailableCopies int32     `json:"availableCopies"`
	Language        string    `json:"language"`
}
