package authenticate

import (
	"context"
	"errors"

	"event_booking_auth_service/internal/config"
	"event_booking_auth_service/internal/domain"
	"event_booking_auth_service/internal/errs"
	"event_booking_auth_service/internal/port/driven"
	"event_booking_auth_service/utils"
)

type UseCase struct {
	cfg        *config.Config
	userStorage driven.UserStorage
}

func New(cfg *config.Config, userStorage driven.UserStorage) *UseCase {
	return &UseCase{
		cfg:        cfg,
		userStorage: userStorage,
	}
}

func (u *UseCase) Authenticate(ctx context.Context, user domain.User) (int, domain.Role, error) {
	userFromDB, err := u.userStorage.GetUserByUserName(ctx, user.UserName)
	if err != nil {
		if !errors.Is(err, errs.ErrNotFound) {
			return 0, "", errs.ErrUserNotFound
		}

		return 0, "", err
	}

	user.Password, err = utils.GenerateHash(user.Password)
	if err != nil {
		return 0, "", err
	}

	if userFromDB.Password != user.Password {
		return 0, "", errs.ErrIncorrectUserNameOrPassword
	}

	return userFromDB.ID, userFromDB.Role, nil
}