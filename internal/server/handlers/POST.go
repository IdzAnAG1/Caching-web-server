package handlers

import (
	"CachingWebServer/internal/domain/models"
	"CachingWebServer/internal/lib/reg"
	"github.com/gin-gonic/gin"
	"time"
)

type registerReq struct {
	AdminToken string `json:"ad_token" binding:"required"`
	Login      string `json:"login" binding:"required,min=8,alphanum"`
	Password   string `json:"password"`
}

func (h *Handlers) RegisterNewUser(c *gin.Context) {
	req := registerReq{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, err.Error())
		return
	}
	if !reg.ValidatePassword(req.Password) {
		c.JSON(400, gin.H{"error": "Password is invalid"})
		return
	}
	models.Users = append(models.Users, models.User{
		ID:           0,
		Login:        req.Login,
		PasswordHash: req.Password,
		CreateAt:     time.Now(),
	})
	c.JSON(200, gin.H{"response": gin.H{"login": req.Login}})
}

func (h *Handlers) GetUsers(c *gin.Context) {
	c.JSON(200, gin.H{"users": models.Users})
}
