package controller

import (
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/dto"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/services/serviceImpl"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type AchievementHandler struct {
	AchievementService serviceImpl.AchievementService
}

func NewAchievementHandler(AchievementService serviceImpl.AchievementService) *AchievementHandler {
	return &AchievementHandler{
		AchievementService: AchievementService,
	}
}

func (h *AchievementHandler) CreateAchievement(c *gin.Context) {
	var createAchievementDTO dto.CreateAchievementDTO
	if err := c.ShouldBindJSON(&createAchievementDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.AchievementService.CreateAchievement(&createAchievementDTO); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create achievement"})
		log.Printf("StartAchievement: failed to add achievement to database: %v", err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "achievement created successfully", "achievement": createAchievementDTO})
}

func (h *AchievementHandler) FindAchievementByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("achievement_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid achievement ID"})
		return
	}

	AchievementDTO, err := h.AchievementService.FindAchievementByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "achievement not found"})
		return
	}

	c.JSON(http.StatusOK, AchievementDTO)
}

func (h *AchievementHandler) FindAllAchievements(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	Achievements, err := h.AchievementService.FindAllAchievements(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve Achievements"})
		return
	}

	c.JSON(http.StatusOK, Achievements)
}

func (h *AchievementHandler) UpdateAchievement(c *gin.Context) {
	var updateAchievementDTO dto.UpdateAchievementDTO
	if err := c.ShouldBindJSON(&updateAchievementDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.AchievementService.UpdateAchievement(&updateAchievementDTO); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Achievement not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Achievement updated successfully", "Achievement": updateAchievementDTO})
}

func (h *AchievementHandler) DeleteAchievement(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("Achievement_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Achievement ID"})
		return
	}

	if err := h.AchievementService.DeleteAchievement(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Achievement not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Achievement deleted successfully"})
}
