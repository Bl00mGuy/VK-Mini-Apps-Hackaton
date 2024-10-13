package controller

import (
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/dto"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/services/impl"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type AchievementHandler struct {
	achievementService impl.AchievementService
}

func NewAchievementHandler(achievementService impl.AchievementService) *AchievementHandler {
	return &AchievementHandler{
		achievementService: achievementService,
	}
}

func (h *AchievementHandler) CreateAchievement(c *gin.Context) {
	var createAchievementDTO dto.CreateAchievementDTO
	if err := c.ShouldBindJSON(&createAchievementDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.achievementService.CreateAchievement(&createAchievementDTO); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create achievement"})
		log.Printf("StartTask: failed to add achievement to database: %v", err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Achievement created successfully", "achievement": createAchievementDTO})
}
