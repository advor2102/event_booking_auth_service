package usecase

import (
	"context"

	"event_booking_auth_service/internal/domain"
)

type Authenticate interface {
	Authenticate(ctx context.Context, user domain.User) (int, domain.Role, error)
}
