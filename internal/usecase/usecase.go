package usecase

import (
	"event_booking_auth_service/internal/adapter/driven/dbstore"
	"event_booking_auth_service/internal/config"
	"event_booking_auth_service/internal/port/usecase"
	empcreator "event_booking_auth_service/internal/usecase/UserCreator"
	authenticate "event_booking_auth_service/internal/usecase/authenticator"
)

type UseCases struct {
	UserCreator   usecase.UserCreator
	Authenticator usecase.Authenticate
}

func New(cfg config.Config, store *dbstore.DBStore) *UseCases {
	return &UseCases{
		UserCreator:   empcreator.New(&cfg, store.UserStorage),
		Authenticator: authenticate.New(&cfg, store.UserStorage),
	}
}
