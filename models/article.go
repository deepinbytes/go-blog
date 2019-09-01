package models

import "github.com/go-ozzo/ozzo-validation"

// Article represents an article record.
type Article struct {
	Id      int    `json:"id" db:"id"`
	Title   string `json:"title" db:"title"`
	Content string `json:"content" db:"content"`
	Author  string `json:"author" db:"author"`
}

// Validate validates the Article fields.
func (m Article) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Title, validation.Required, validation.Length(0, 120)),
		validation.Field(&m.Content, validation.Required, validation.Length(0, 6000)),
		validation.Field(&m.Author, validation.Required, validation.Length(0, 120)),
	)
}
