package controller

import (
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/dto"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/services/serviceImpl"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type AvatarHandler struct {
	AvatarService serviceImpl.AvatarService
}

func NewAvatarHandler(AvatarService serviceImpl.AvatarService) *AvatarHandler {
	return &AvatarHandler{
		AvatarService: AvatarService,
	}
}

func (h *AvatarHandler) CreateAvatar(c *gin.Context) {
	var createAvatarDTO dto.CreateAvatarDTO
	if err := c.ShouldBindJSON(&createAvatarDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.AvatarService.CreateAvatar(&createAvatarDTO); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create avatar"})
		log.Printf("StartAvatar: failed to add avatar to database: %v", err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Avatar created successfully", "Avatar": createAvatarDTO})
}

func (h *AvatarHandler) FindAvatarByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("Avatar_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Avatar ID"})
		return
	}

	AvatarDTO, err := h.AvatarService.FindAvatarByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Avatar not found"})
		return
	}

	c.JSON(http.StatusOK, AvatarDTO)
}

func (h *AvatarHandler) FindAllAvatars(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	Avatars, err := h.AvatarService.FindAllAvatars(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve Avatars"})
		return
	}

	c.JSON(http.StatusOK, Avatars)
}

func (h *AvatarHandler) UpdateAvatar(c *gin.Context) {
	var updateAvatarDTO dto.UpdateAvatarDTO
	if err := c.ShouldBindJSON(&updateAvatarDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.AvatarService.UpdateAvatar(&updateAvatarDTO); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Avatar not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Avatar updated successfully", "Avatar": updateAvatarDTO})
}

func (h *AvatarHandler) DeleteAvatar(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("avatar_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Avatar ID"})
		return
	}

	if err := h.AvatarService.DeleteAvatar(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Avatar not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Avatar deleted successfully"})
}
