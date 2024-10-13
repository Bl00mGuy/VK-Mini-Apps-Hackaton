package mapper

import (
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/dto"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/entity"
)

type AchievementMapper struct {
}

func NewAchievementMapper() *AchievementMapper {
	return &AchievementMapper{}
}

func (a *AchievementMapper) ConvertToDTO(achievement *entity.Achievement) *dto.AchievementDTO {
	return &dto.AchievementDTO{
		AchievementID: achievement.ID,
		UserID:        achievement.UserID,
		Title:         achievement.Title,
		Description:   achievement.Description,
	}
}

func (a *AchievementMapper) ConvertToSliceDTO(achievements []entity.Achievement) []dto.AchievementDTO {
	arr := make([]dto.AchievementDTO, 0, len(achievements))
	for _, achievement := range achievements {
		achievementDTO := a.ConvertToDTO(&achievement)
		arr = append(arr, *achievementDTO)
	}
	return arr
}

func ConvertToAchievementDTO(achievement entity.Achievement) dto.AchievementDTO {
	return dto.AchievementDTO{
		AchievementID: achievement.ID,
		UserID:        achievement.UserID,
		Title:         achievement.Title,
		Description:   achievement.Description,
	}
}
