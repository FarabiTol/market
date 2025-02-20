package postgres

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"market/src/domain"
	"market/src/repository"
)

type marketRepository struct {
	db     *pgxpool.Pool
	logger log.Logger
}

func (c *marketRepository) CreateCollection(ctx context.Context, collection *domain.Collection) error {
	_ = level.Info(c.logger).Log("msg", fmt.Sprintf("...Inserting company information: %+v", collection))

	conn, err := c.db.Acquire(ctx)
	if err != nil {
		_ = level.Error(c.logger).Log(err)
		return err
	}

	defer conn.Release()

	sqlStatement := `INSERT INTO collection (id, user_id, short_name, full_name, status, ammount, currency, register_date, created_at, updated_at)
	        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`

	_, err = conn.Exec(ctx, sqlStatement,
		collection.ID,
		collection.UserID,
		collection.ShortName,
		collection.FullName,
		collection.Status,
		collection.Amount,
		collection.Currency,
		collection.RegisterDate,
		collection.CreatedAt,
		collection.UpdatedAt,
	)
	if err != nil {
		_ = level.Error(c.logger).Log("msg", "Error while inserting company info", "error", err.Error())

		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == UniqueViolationErrorCode {
			return repository.ErrAlreadyExist
		}

		return err
	}

	return nil
}
func (c *marketRepository) GetCollection(ctx context.Context, id uuid.UUID) (*domain.Collection, error) {
	conn, err := c.db.Acquire(ctx)
	if err != nil {
		_ = level.Error(c.logger).Log(err)
		return nil, err
	}

	defer conn.Release()

	var collection domain.Collection

	sqlStatement := `SELECT id, user_id, short_name, full_name, status, ammount, currency, register_date, created_at, updated_at FROM collection WHERE id=$1`

	row := conn.QueryRow(ctx, sqlStatement, id)
	err = row.Scan(
		&collection.ID,
		&collection.UserID,
		&collection.ShortName,
		&collection.FullName,
		&collection.Status,
		&collection.Amount,
		&collection.Currency,
		&collection.RegisterDate,
		&collection.CreatedAt,
		&collection.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("NotFound")
		}
		return nil, err
	}

	return &collection, nil
}
