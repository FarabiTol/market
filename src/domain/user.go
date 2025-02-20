package domain

import "github.com/google/uuid"

type User struct {
	ID          uuid.UUID `json:"id"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phoneNumber"`
	AccountID   uuid.UUID `json:"accountID"`
}
