package repository

import (
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/dto"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/entity"
)

type AchievementRepository interface {
	Create(createTaskDTO *dto.CreateAchievementDTO) error
	FindByID(findTaskDTO *dto.FindAchievementDTO) (*entity.Achievement, error)
	FindAll(findAllTasksDTO *dto.FindAllAchievementsDTO) ([]entity.Achievement, error)
	Update(updateTaskDTO *dto.UpdateAchievementDTO) error
	Delete(deleteTaskDTO *dto.DeleteAchievementDTO) error
}
