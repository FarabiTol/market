package domain

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"time"
)

type Collection struct {
	ID           uuid.UUID       `json:"id"`
	UserID       uuid.UUID       `json:"userID"`
	ShortName    string          `json:"shortName"`
	FullName     string          `json:"fullName"`
	Status       string          `json:"status"`
	Amount       decimal.Decimal `json:"amount"`
	Currency     string          `json:"currency"`
	RegisterDate time.Time       `json:"registerDate"`
	UpdatedAt    time.Time       `json:"updatedAt"`
	CreatedAt    time.Time       `json:"createdAt"`
}
