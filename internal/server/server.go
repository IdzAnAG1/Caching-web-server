package server

import (
	"CachingWebServer/internal/config"
	"CachingWebServer/internal/server/handlers"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type Server struct {
	Port       string
	Token      string
	TokenTTL   time.Duration
	AdminToken string
}

func New(cfg config.Config) *Server {
	return &Server{
		Port:       strconv.Itoa(cfg.Server.Port),
		Token:      cfg.Server.Token,
		TokenTTL:   cfg.Server.TokenTTL,
		AdminToken: cfg.Admin.Token,
	}
}
func (s *Server) Launch() error {
	H := handlers.NewHandlers(s.Token, s.AdminToken, s.TokenTTL)

	r := gin.Default()
	r.POST("/api/register", H.RegisterNewUser)
	r.GET("/api/users", H.GetUsers)
	r.POST("/login", H.Login)
	err := r.Run("localhost:" + s.Port)
	if err != nil {
		return err
	}
	return nil
}
