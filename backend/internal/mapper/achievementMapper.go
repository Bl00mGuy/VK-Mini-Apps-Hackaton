package mapper

import (
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/dto"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/entity"
)

func convertToAchievementDTO(achievement entity.Achievement) dto.AchievementDTO {
	return dto.AchievementDTO{
		AchievementID: achievement.ID,
		UserID:        achievement.UserID,
		Title:         achievement.Title,
		Description:   achievement.Description,
	}
}
