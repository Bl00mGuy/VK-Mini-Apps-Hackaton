package repository

import (
 "github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/dto"
)

type TaskRepository interface {
 Create(createTaskDTO *dto.CreateTaskDTO) error
 FindByID(id uint) (*dto.TaskDTO, error)
 FindAll(userID uint) ([]dto.TaskDTO, error)
 Update(updateTaskDTO *dto.UpdateTaskDTO) error
 Delete(id uint) error
}
