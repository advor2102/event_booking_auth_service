package usecase

import (
	"context"

	"event_booking_auth_service/internal/domain"
)

type UserCreator interface {
	CreateUser(ctx context.Context, user domain.User) (err error)
}
