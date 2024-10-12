package repository

import (
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/dto"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/entity"
)

type TaskRepository interface {
	Create(createTaskDTO *dto.CreateTaskDTO) error
	FindByID(findTaskDTO *dto.FindTaskDTO) (*entity.Task, error)
	FindAll(findAllTasksDTO *dto.FindAllTasksDTO) ([]entity.Task, error)
	Update(updateTaskDTO *dto.UpdateTaskDTO) error
	Delete(deleteTaskDTO *dto.DeleteTaskDTO) error
}
