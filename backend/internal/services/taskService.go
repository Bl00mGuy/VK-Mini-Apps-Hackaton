package services

import "github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/dto"

type TaskServiceInterface interface {
	CreateTask(createTaskDTO *dto.CreateTaskDTO) error
	FindTaskByID(id uint) (*dto.TaskDTO, error)
	FindAllTasks(userID uint) ([]dto.TaskDTO, error)
	UpdateTask(updateTaskDTO *dto.UpdateTaskDTO) error
	DeleteTask(id uint) error
}
