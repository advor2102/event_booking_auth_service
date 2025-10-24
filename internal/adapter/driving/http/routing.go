package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) endpoints() {
	s.router.GET("/ping", s.Ping)

	authG := s.router.Group("/auth")
	authG.POST("/sign-up", s.SignUp)
	authG.POST("/sign-in", s.SignIn)
	authG.GET("/refresh", s.RefreshTokenPairs)
}

func (s *Server) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ping": "pong",
	})
}
