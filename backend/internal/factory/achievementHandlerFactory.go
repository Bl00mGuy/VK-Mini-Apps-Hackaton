package factory

import (
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/controller"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/mapper"
	impl2 "github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/repository/impl"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/services/impl"
	"gorm.io/gorm"
)

func CreateAchievementHandler(db *gorm.DB) *controller.AchievementHandler {
	achievementMapper := mapper.NewAchievementMapper()
	achievementRepository := impl2.NewAchievementRepository(db, achievementMapper)
	achievementService := impl.NewAchievementService(achievementRepository)

	return controller.NewAchievementHandler(*achievementService)
}
