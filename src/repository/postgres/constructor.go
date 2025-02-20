package postgres

import (
	"context"
	"github.com/go-kit/log"
	"github.com/jackc/pgx/v4/pgxpool"
	"market/src/config"
	"market/src/repository"
)

const UniqueViolationErrorCode = "23505"

const dbType = "postgresql"

type Database struct {
	db     *pgxpool.Pool
	logger log.Logger

	marketRepo repository.MarketRepository
}

func New(cfg config.PostgresConfig, logger log.Logger) (repository.Storage, error) {
	db, err := postgresql.InitConnect(
		context.Background(),
		cfg.MaxConns,
		cfg.Host,
		cfg.Port,
		cfg.Username,
		cfg.Password,
		cfg.DatabaseName,
	)
	if err != nil {
		return nil, err
	}

	store := &Database{
		db:     db,
		logger: log.With(logger, "module", "postgres"),
	}

	if err = store.migrateUp(cfg); err != nil {
		return nil, err
	}

	return store, nil
}

func (d *Database) Close() error {
	d.db.Close()
	return nil
}

func (d *Database) Name() string {
	return "postgres"
}

func (d *Database) MarketRepository() repository.MarketRepository {
	if d.marketRepo == nil {
		d.marketRepo = &marketRepository{
			db:     d.db,
			logger: log.With(d.logger, "repository", "market"),
		}
	}

	return d.marketRepo
}
