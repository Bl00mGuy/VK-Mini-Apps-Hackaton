package factory

import (
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/controller"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/mapper"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/repository/repoImpl"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/services/serviceImpl"
	"gorm.io/gorm"
)

func CreateUserHandler(db *gorm.DB) *controller.UserHandler {
	UserMapper := mapper.NewUserMapper()
	UserRepository := repoImpl.NewUserRepository(db, UserMapper)
	UserService := serviceImpl.NewUserService(UserRepository)

	return controller.NewUserHandler(*UserService)
}
