package handlers

import (
	"CachingWebServer/internal/domain/models"
	"CachingWebServer/internal/lib/crypt"
	"CachingWebServer/internal/lib/jwt"
	"CachingWebServer/internal/lib/reg"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"os"
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
	if req.AdminToken != h.AdminToken {
		c.JSON(400, gin.H{"Error": "Access denied"})
	}
	if !reg.ValidatePassword(req.Password) {
		c.JSON(400, gin.H{"error": "Password is invalid"})
		return
	}
	hashedPass, err := crypt.Hash(req.Password)
	if err != nil {
		c.JSON(400, gin.H{"error": "Internal Error"})
	}
	user := models.User{
		ID:           uuid.NewString(),
		Login:        req.Login,
		PasswordHash: hashedPass,
		CreateAt:     time.Time{},
	}
	// Временное решение пока не прикрючена БД
	models.Users[req.Login] = user
	c.JSON(200, gin.H{"response": gin.H{"login": req.Login}})
}

type LoginReq struct {
	Login    string `json:"login" binding:"required,min=8,alphanum"`
	Password string `json:"password" `
}

func (h *Handlers) Login(c *gin.Context) {
	req := LoginReq{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(200, err.Error())
	}
	if !reg.ValidatePassword(req.Password) {
		c.JSON(400, gin.H{"Error": "Password is invalid"})
		return
	}
	if !crypt.HashCompare(models.Users[req.Login].PasswordHash, req.Password) {
		c.JSON(400, gin.H{"Error": "Invalid password"})
		return
	}
	token, err := jwt.SignToken(models.Users[req.Login].ID, h.ServerToken, h.TTL)
	if err != nil {
		c.JSON(400, gin.H{"Internal Error": fmt.Sprintf("token err :%v", err.Error())})
		return
	}
	c.JSON(200, gin.H{"resp": token})
}

type UploadReq struct {
	Name string  `json:"Name" binding:"required"`
	file os.File `json:"File"`
}

func (h *Handlers) UploadTheFile(c *gin.Context) {
	req := UploadReq{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"Internal Error": "Binding Req"})
		return
	}
	file, err := c.FormFile(req.Name)
	if err != nil {
		c.JSON(400, gin.H{"error": "File is required"})
		return
	}
	if err := c.SaveUploadedFile(file, "./TEMP"); err != nil {
		c.JSON(400, gin.H{"Internal Error": "Failed attempt to save a file"})
		return
	}
	c.JSON(200, gin.H{"resp": "NORM"})
}
