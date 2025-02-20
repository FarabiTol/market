package service

import (
	"context"
	"github.com/go-kit/log"
	"github.com/google/uuid"
	"market/src/api/transport"
	"market/src/domain"
	"market/src/repository"
	"time"
)

type marketService struct {
	store  repository.Storage
	logger log.Logger
}

func NewMarketService(store repository.Storage, logger log.Logger) MarketService {
	return &marketService{
		store:  store,
		logger: logger,
	}
}

func (m *marketService) CreateCollection(ctx context.Context, collection *transport.CreateCollectionReq) (uuid.UUID, error) {
	id := uuid.New()

	coll := &domain.Collection{
		ID:           id,
		UserID:       collection.UserID,
		ShortName:    collection.ShortName,
		FullName:     collection.FullName,
		Status:       collection.Status,
		Amount:       collection.Amount,
		Currency:     collection.Currency,
		RegisterDate: collection.RegisterDate,
		UpdatedAt:    time.Now(),
		CreatedAt:    time.Now(),
	}

	err := m.store.MarketRepository().CreateCollection(ctx, coll)
	if err != nil {
		return uuid.UUID{}, err
	}

	return id, nil
}

func (m *marketService) GetCollection(ctx context.Context, id uuid.UUID) (*domain.Collection, error) {
	coll, err := m.store.MarketRepository().GetCollection(ctx, id)
	if err != nil {
		return nil, err
	}

	return coll, nil
}
