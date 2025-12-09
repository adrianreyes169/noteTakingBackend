package note

import "time"

type Note struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
