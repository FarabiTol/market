package repository

type Storage interface {
	Name() string
	Close() error

	MarketRepository() MarketRepository
}
