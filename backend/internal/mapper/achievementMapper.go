package mapper

import (
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/dto"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/entity"
)

type AchievementMapper struct{}

func NewAchievementMapper() *AchievementMapper {
	return &AchievementMapper{}
}

func (a *AchievementMapper) ConvertToAchievementDTOs(achievements []entity.Achievement) []dto.AchievementDTO {
	var achievementDTOs []dto.AchievementDTO

	for i, achievement := range achievements {
		achievementDTOs[i] = a.ConvertToAchievementDTO(achievement)
	}
	return achievementDTOs
}

func (a *AchievementMapper) ConvertToAchievementDTO(achievement entity.Achievement) dto.AchievementDTO {
	return dto.AchievementDTO{
		AchievementID: achievement.ID,
		UserID:        achievement.UserID,
		Title:         achievement.Title,
		Description:   achievement.Description,
	}
}
