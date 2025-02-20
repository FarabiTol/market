package transport

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"time"
)

type CreateCollectionReq struct {
	UserID       uuid.UUID       `json:"userID"`
	ShortName    string          `json:"shortName"`
	FullName     string          `json:"fullName"`
	Status       string          `json:"status"`
	Amount       decimal.Decimal `json:"amount"`
	Currency     string          `json:"currency"`
	RegisterDate time.Time       `json:"registerDate"`
}
