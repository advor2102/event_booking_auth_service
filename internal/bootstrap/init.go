package bootstrap

import (
	"net/http"

	http2 "event_booking_auth_service/internal/adapter/driving/http"
	"event_booking_auth_service/internal/config"
	"event_booking_auth_service/internal/usecase"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func initDB(cfg config.Postgres, name string) (*sqlx.DB, error) {
	connConfig, err := pgx.ParseConfig(cfg.ConnectionURL())
	if err != nil {
		return nil, err
	}

	connStr := stdlib.RegisterConnConfig(connConfig)

	db, err := sqlx.Connect("pgx", connStr)
	if err != nil {
		return db, err
	}

	db.SetMaxOpenConns(cfg.MaxOpenConnections)
	db.SetMaxIdleConns(cfg.MaxIdleConnections)
	db.SetConnMaxLifetime(cfg.ConnectionMaxLifetime)

	return db, nil
}

func initHTTPService(cfg *config.Config, uc *usecase.UseCases) *http.Server {
	return http2.New(cfg, uc)
}
