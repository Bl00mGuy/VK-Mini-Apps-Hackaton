package factory

import (
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/controller"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/mapper"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/repository/repoImpl"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/services/serviceImpl"
	"gorm.io/gorm"
)

func CreateTaskHandler(db *gorm.DB) *controller.TaskHandler {
	taskMapper := mapper.NewTaskMapper()
	taskRepository := repoImpl.NewTaskRepository(db, taskMapper)
	taskService := serviceImpl.NewTaskService(taskRepository)

	return controller.NewTaskHandler(*taskService)
}
