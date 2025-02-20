package postgres

import (
	"errors"
	"github.com/golang-migrate/migrate/v4"
	"market/src/config"
)

import (
	"github.com/go-kit/log/level"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func (d *Database) migrateUp(cfg config.PostgresConfig) error {
	dbURL := "postgres://" +
		cfg.Username + ":" +
		cfg.Password + "@" +
		cfg.Host + ":" +
		cfg.Port + "/" +
		cfg.DatabaseName +
		"?sslmode=disable"

	m, err := migrate.New(
		"file://schema/migration",
		dbURL,
	)
	if err != nil {
		_ = level.Error(d.logger).Log("migration error", err)
		return err
	}

	err = m.Up()
	if err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			_ = level.Info(d.logger).Log("migration status", "no change")
			return nil // No error, just no change needed
		}
		_ = level.Error(d.logger).Log("migration error", err)
		return err
	}

	return nil
}
