package driven

import (
	"context"

	"event_booking_auth_service/internal/domain"
)

type UserStorage interface {
	CreateUser(ctx context.Context, user domain.User) (err error)
	GetUserByID(ctx context.Context, id int) (domain.User, error)
	GetUserByUserName(ctx context.Context, userName string) (domain.User, error)
}