package factory

import (
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/controller"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/mapper"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/repository"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/services/impl"
	"gorm.io/gorm"
)

func CreateTaskHandler(db *gorm.DB) *controller.TaskHandler {
	taskMapper := mapper.NewTaskMapper()
	taskRepository := repository.NewTaskRepository(db, taskMapper)
	taskService := impl.NewTaskService(taskRepository)

	return controller.NewTaskHandler(*taskService)
}
