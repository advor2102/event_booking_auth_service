package empcreator

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

func New(cfg *config.Config, userStorage driven.UserStorage) *UseCase{
	return &UseCase{
		cfg: cfg,
		userStorage: userStorage,
	}
}

func (u *UseCase) CreateUser(ctx context.Context, user domain.User) (err error) {
	_, err = u.userStorage.GetUserByUserName(ctx, user.UserName)
	if err != nil {
		if !errors.Is(err, errs.ErrNotFound) {
			return err
		}
	} else {
		return errs.ErrUserNameAlreadyExist
	}

	user.Password, err = utils.GenerateHash(user.Password)
	if err != nil {
		return err
	}

	user.Role = domain.RoleUser

	if err = u.userStorage.CreateUser(ctx, user); err != nil {
		return err
	}

	return nil
}