package repository

import (
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/dto"
)

type TaskRepository interface {
	Create(createTaskDTO *dto.CreateTaskDTO) error
	FindByID(findTaskDTO *dto.FindTaskDTO) (*dto.TaskDTO, error)
	FindAll(findAllTasksDTO *dto.FindAllTasksDTO) ([]dto.TaskDTO, error)
	Update(updateTaskDTO *dto.UpdateTaskDTO) error
	Delete(deleteTaskDTO *dto.DeleteTaskDTO) error
}
