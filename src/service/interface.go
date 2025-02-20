package service

import (
	"context"
	"github.com/google/uuid"
	"market/src/api/transport"
	"market/src/domain"
)

type MarketService interface {
	CreateCollection(ctx context.Context, collection *transport.CreateCollectionReq) (uuid.UUID, error)
	GetCollection(ctx context.Context, id uuid.UUID) (*domain.Collection, error)
}
