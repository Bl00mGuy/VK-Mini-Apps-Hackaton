package serviceImpl

import (
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/dto"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/repository"
)

type TaskService struct {
	repo repository.TaskRepository
}

func NewTaskService(taskRepository repository.TaskRepository) *TaskService {
	return &TaskService{
		repo: taskRepository,
	}
}

func (s *TaskService) CreateTask(createTaskDTO *dto.CreateTaskDTO) error {
	return s.repo.Create(createTaskDTO)
}

func (s *TaskService) FindTaskByID(id uint) (*dto.TaskDTO, error) {
	return s.repo.FindByID(id)
}

func (s *TaskService) FindAllTasks(userID uint) ([]dto.TaskDTO, error) {
	return s.repo.FindAll(userID)
}

func (s *TaskService) UpdateTask(updateTaskDTO *dto.UpdateTaskDTO) error {
	return s.repo.Update(updateTaskDTO)
}

func (s *TaskService) DeleteTask(id uint) error {
	return s.repo.Delete(id)
}
