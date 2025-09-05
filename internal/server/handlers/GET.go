package handlers

import (
	"CachingWebServer/internal/domain/models"
	"github.com/gin-gonic/gin"
)

type GetHandlers struct {
	H Handlers
}

func (h *Handlers) GetUsers(c *gin.Context) {
	c.JSON(200, gin.H{"users": models.Users})
}

func (h *Handlers) GetUserFiles(c *gin.Context) {

}
