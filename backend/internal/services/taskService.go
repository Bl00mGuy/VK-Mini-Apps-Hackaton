package services

import "github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/dto"

type TaskServiceInterface interface {
	CreateTask(createTaskDTO *dto.CreateTaskDTO) error
	FindTaskByID(findTaskDTO *dto.FindTaskDTO) (*dto.TaskDTO, error)
	FindAllTasks(findAllTasksDTO *dto.FindAllTasksDTO) ([]dto.TaskDTO, error)
	UpdateTask(updateTaskDTO *dto.UpdateTaskDTO) error
	DeleteTask(deleteTaskDTO *dto.DeleteTaskDTO) error
}
