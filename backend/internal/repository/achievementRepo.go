package repository

import (
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/dto"
)

type AchievementRepository interface {
	Create(createTaskDTO *dto.CreateAchievementDTO) error
	FindByID(achievementID uint) (*dto.AchievementDTO, error)
	FindAll(userID uint) ([]dto.AchievementDTO, error)
	Update(updateTaskDTO *dto.UpdateAchievementDTO) error
	Delete(achievementID uint) error
}
