package http

import (
	"errors"
	"net/http"

	"event_booking_auth_service/internal/domain"
	"event_booking_auth_service/internal/errs"
	"event_booking_auth_service/pkg"

	"github.com/gin-gonic/gin"
)

type SignUpRequest struct {
	FullName string `json:"full_name" db:"full_name"`
	UserName string `json:"user_name" db:"user_name"`
	Password string `json:"password" db:"password"`
}

// SignUp
// @Summary Create user
// @Description Create new user and add to database
// @Tags Auth
// @Consume json
// @Produce json
// @Param request_body body SignUpRequest true "new euser data"
// @Success 201 {object} CommonResponse
// @Failure 400 {object} CommonError
// @Failure 404 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /auth/sign-up [post]
func (s *Server) SignUp(c *gin.Context) {
	var input SignUpRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		s.handleError(c, errors.Join(errs.ErrInvalidRequestBody, err))
		return
	}

	if err := s.uc.UserCreator.CreateUser(c, domain.User{
		FullName: input.FullName,
		UserName: input.UserName,
		Password: input.Password,
	}); err != nil {
		s.handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, CommonResponse{Message: "User created successfully"})
}

type SignInRequest struct {
	UserName string `json:"user_name" db:"user_name"`
	Password string `json:"password" db:"password"`
}

type TokenPairResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// SignIn
// @Summary Enter
// @Description Enter as user
// @Tags Auth
// @Consume json
// @Produce json
// @Param request_body body SignIpRequest true "login and password"
// @Success 200 {object} TokenPairResponse
// @Failure 400 {object} CommonError
// @Failure 404 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /auth/sign-in [post]
func (s *Server) SignIn(c *gin.Context) {
	var input SignInRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		s.handleError(c, errors.Join(errs.ErrInvalidRequestBody, err))
		return
	}

	employeeID, employeeRole, err := s.uc.Authenticator.Authenticate(c, domain.User{
		UserName: input.UserName,
		Password: input.Password,
	})
	if err != nil {
		s.handleError(c, err)
		return
	}

	accessToken, refreshToken, err := s.generateNewTokenPair(employeeID, employeeRole)
	if err != nil {
		s.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, TokenPairResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

const (
	refreshTokenHeader = "X-Refresh_token"
)

// RefreshTokenPairs
// @Summary Refresh token pairs
// @Description Refresh token pairs
// @Tags Auth
// @Produce json
// @Param X-Refresh_token header string true "input refresh token"
// @Success 200 {object} TokenPairResponse
// @Failure 400 {object} CommonError
// @Failure 404 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /auth/refresh [get]
func (s *Server) RefreshTokenPairs(c *gin.Context) {
	token, err := s.extractTokenFromHeader(c, refreshTokenHeader)
	if err != nil {
		c.JSON(http.StatusUnauthorized, CommonError{Error: err.Error()})
		return
	}

	employeeID, isRefresh, employeeRole, err := pkg.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, CommonError{Error: err.Error()})
		return
	}

	if !isRefresh {
		c.JSON(http.StatusUnauthorized, CommonError{Error: "inappropriate token"})
		return
	}

	accessToken, refreshToken, err := s.generateNewTokenPair(employeeID, employeeRole)
	if err != nil {
		s.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, TokenPairResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}
