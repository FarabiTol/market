package repository

import (
	"context"
	"github.com/google/uuid"
	"market/src/domain"
)

type MarketRepository interface {
	CreateCollection(ctx context.Context, collection *domain.Collection) error
	GetCollection(ctx context.Context, id uuid.UUID) (*domain.Collection, error)
}
