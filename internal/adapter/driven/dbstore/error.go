package dbstore

import (
	"database/sql"
	"errors"

	"event_booking_auth_service/internal/errs"
)

func (u *UserStorage) translateError(err error) error {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return errs.ErrNotFound
	default:
		return err
	}
}
