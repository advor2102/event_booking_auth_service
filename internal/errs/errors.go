package errs

import "errors"

var (
	ErrUserNotFound                = errors.New("user not found")
	ErrInvalidUserID               = errors.New("invalid user id")
	ErrNotFound                    = errors.New("not found")
	ErrInvalidRequestBody          = errors.New("invalid request body")
	ErrInvalidFieldValue           = errors.New("invalid field value")
	ErrUserNameAlreadyExist        = errors.New("user name already exist")
	ErrIncorrectUserNameOrPassword = errors.New("incorrect employee name or password")
	ErrInvalidToken                = errors.New("invalid token")
	ErrSomethingWentWrong          = errors.New("something went wrong")
)
