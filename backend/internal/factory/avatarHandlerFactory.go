package factory

import (
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/controller"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/mapper"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/repository/repoImpl"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/services/serviceImpl"
	"gorm.io/gorm"
)

func CreateAvatarHandler(db *gorm.DB) *controller.AvatarHandler {
	AvatarMapper := mapper.NewAvatarMapper()
	AvatarRepository := repoImpl.NewAvatarRepository(db, AvatarMapper)
	AvatarService := serviceImpl.NewAvatarService(AvatarRepository)

	return controller.NewAvatarHandler(*AvatarService)
}
