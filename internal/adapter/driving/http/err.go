package http

import (
	"errors"
	"net/http"

	"event_booking_auth_service/internal/errs"
	"github.com/gin-gonic/gin"
)

func (s *Server) handleError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, errs.ErrUserNotFound) || errors.Is(err, errs.ErrNotFound):
		c.JSON(http.StatusNotFound, CommonError{Error: err.Error()})
	case errors.Is(err, errs.ErrInvalidUserID) || errors.Is(err, errs.ErrInvalidRequestBody):
		c.JSON(http.StatusBadRequest, CommonError{Error: err.Error()})
	case errors.Is(err, errs.ErrInvalidFieldValue) || errors.Is(err, errs.ErrUserNameAlreadyExist):
		c.JSON(http.StatusUnprocessableEntity, CommonError{Error: err.Error()})
	case errors.Is(err, errs.ErrIncorrectUserNameOrPassword) || errors.Is(err, errs.ErrInvalidToken):
		c.JSON(http.StatusUnauthorized, CommonError{Error: err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, CommonError{Error: err.Error()})
	}
}
