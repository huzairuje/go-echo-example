package entity

import (
	"github.com/gofrs/uuid"
	"time"
)

type Product struct {
	Id          uuid.UUID `json:"id" gorm:"primary_key; unique;type:uuid; column:id;default:uuid_generate_v4()"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Rating      int       `json:"rating"`
	Image       string    `json:"image"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}
