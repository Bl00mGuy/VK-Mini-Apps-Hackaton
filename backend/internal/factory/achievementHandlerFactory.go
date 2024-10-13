package factory

import (
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/controller"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/mapper"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/repository/repoImpl"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/services/serviceImpl"
	"gorm.io/gorm"
)

func CreateAchievementHandler(db *gorm.DB) *controller.AchievementHandler {
	achievementMapper := mapper.NewAchievementMapper()
	achievementRepository := repoImpl.NewAchievementRepository(db, achievementMapper)
	achievementService := serviceImpl.NewAchievementService(achievementRepository)

	return controller.NewAchievementHandler(*achievementService)
}
