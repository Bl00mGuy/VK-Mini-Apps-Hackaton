package repository

import (
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/dto"
)

type AchievementRepository interface {
	Create(createAchievementDTO *dto.CreateAchievementDTO) error
	FindByID(id uint) (*dto.AchievementDTO, error)
	FindAll(userID uint) ([]dto.AchievementDTO, error)
	Update(updateAchievementDTO *dto.UpdateAchievementDTO) error
	Delete(id uint) error
	LinkUserAchievement(userID uint, achievementID uint) error
}
