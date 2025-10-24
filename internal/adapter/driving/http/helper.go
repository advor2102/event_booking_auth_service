package http

import (
	"errors"
	"strings"

	"event_booking_auth_service/internal/domain"
	"event_booking_auth_service/pkg"
	"github.com/gin-gonic/gin"
)

func (s *Server) extractTokenFromHeader(c *gin.Context, headerKey string) (string, error) {
	header := c.GetHeader(headerKey)

	if header == "" {
		return "", errors.New("empty authorization header")
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		return "", errors.New("invalid authorization header")
	}
	if len(headerParts[1]) == 0 {
		return "", errors.New("empty token")
	}
	return headerParts[1], nil
}

func (s *Server) generateNewTokenPair(employeeID int, employeeRole domain.Role) (string, string, error) {
	accessToken, err := pkg.GenerateToken(employeeID, s.cfg.AuthParams.AccessTokenTtlMinutes, employeeRole, false)
	if err != nil {
		return "", "", err
	}
	refreshToken, err := pkg.GenerateToken(employeeID, s.cfg.AuthParams.RefreshTokenTtlDays, employeeRole, true)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}