package services

import "github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/dto"

type AchievementServiceInterface interface {
	CreateAchievement(createAchievementDTO *dto.CreateAchievementDTO) error
	FindAchievementByID(id uint) (*dto.AchievementDTO, error)
	FindAllAchievements(userID uint) ([]dto.AchievementDTO, error)
	UpdateAchievement(updateAchievementDTO *dto.UpdateAchievementDTO) error
	DeleteAchievement(id uint) error
}
