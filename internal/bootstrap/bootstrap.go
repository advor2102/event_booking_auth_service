package bootstrap

import (
	"context"
	"fmt"

	"event_booking_auth_service/internal/adapter/driven/dbstore"
	"event_booking_auth_service/internal/config"
	"event_booking_auth_service/internal/usecase"
)

func initLayers(cfg config.Config) *App {
	teardown := make([]func(), 0)

	db, err := initDB(*cfg.Postgres, "hpay_astrasend")
	if err != nil {
		panic(err)
	}

	storage := dbstore.New(db)

	teardown = append(teardown, func() {
		if err := db.Close(); err != nil {
			fmt.Println(err)
		}
	})

	uc := usecase.New(cfg, storage)

	httpSrv := initHTTPService(&cfg, uc)

	teardown = append(teardown, func() {
		ctxShutDown, cancel := context.WithTimeout(context.Background(), gracefulDeadline)
		defer cancel()
		if err := httpSrv.Shutdown(ctxShutDown); err != nil {
			return
		}
	})

	return &App{
		cfg:      cfg,
		rest:     httpSrv,
		teardown: teardown,
	}
}
