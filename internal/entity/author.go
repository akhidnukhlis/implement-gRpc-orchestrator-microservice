package entity

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Author struct {
	ID        string    `gorm:"size:36;not null;unique index;primaryKey"`
	Name      string    `gorm:"size:255;"`
	Nickname  string    `gorm:"size:100;unique"`
	Email     string    `gorm:"size:100;unique"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

// Authors represent body for get data from author
type Authors struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Nickname  string `json:"nickname"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// AuthorRequest is payload for register author
type AuthorRequest struct {
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
}

// FindAuthorRequest is payload for get author
type FindAuthorRequest struct {
	ID string `json:"id"`
}

// AuthorRequestValidate is to validate input request
func AuthorRequestValidate(ar *AuthorRequest) error {
	err := validation.Errors{
		"name":     validation.Validate(&ar.Name, validation.Required, validation.Length(2, 40)),
		"nickname": validation.Validate(&ar.Nickname, validation.Required),
		"email":    validation.Validate(&ar.Email, validation.Required),
	}

	return err.Filter()
}
