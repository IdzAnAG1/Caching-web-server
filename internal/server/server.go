package server

import (
	"CachingWebServer/internal/config"
	"CachingWebServer/internal/server/handlers"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type Server struct {
	port     string
	adToken  string
	TokenTTL time.Duration
}

func New(cfg config.Config) *Server {
	return &Server{
		port:     strconv.Itoa(cfg.Server.Port),
		adToken:  cfg.Server.AdminToken,
		TokenTTL: cfg.Server.TokenTTL,
	}
}
func (s *Server) Launch() error {
	H := handlers.NewHandlers(s.adToken)

	r := gin.Default()
	r.POST("/api/register", H.RegisterNewUser)
	r.GET("/api/users", H.GetUsers)
	r.POST("/login", H.Login)
	err := r.Run("localhost:" + s.port)
	if err != nil {
		return err
	}
	return nil
}
