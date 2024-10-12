package mapper

import (
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/dto"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/entity"
)

type TaskMapper struct {
}

func NewTaskMapper() *TaskMapper {
	return &TaskMapper{}
}

func (t *TaskMapper) ConvertToObject(dto *dto.UpdateTaskDTO) *entity.Task {
	return &entity.Task{
		ID:          dto.TaskID,
		UserID:      dto.UserID,
		Title:       dto.Title,
		Description: dto.Description,
		IsCompleted: dto.IsCompleted,
	}
}

func (t *TaskMapper) ConvertToDTO(task *entity.Task) *dto.TaskDTO {
	return &dto.TaskDTO{
		TaskID:      task.ID,
		UserID:      task.UserID,
		Title:       task.Title,
		Description: task.Description,
		IsCompleted: task.IsCompleted,
	}
}

func (t *TaskMapper) ConvertToSliceDTO(tasks []entity.Task) []dto.TaskDTO {
	arr := make([]dto.TaskDTO, len(tasks))
	for i := range tasks {
		taskDTO := dto.TaskDTO{
			TaskID:      tasks[i].ID,
			UserID:      tasks[i].UserID,
			Title:       tasks[i].Title,
			Description: tasks[i].Description,
			IsCompleted: tasks[i].IsCompleted,
		}
		arr = append(arr, taskDTO)
	}
	return arr
}

func (t *TaskMapper) ConvertToSliceObjects(tasks []dto.TaskDTO) []entity.Task {
	arr := make([]entity.Task, len(tasks))
	for i := range tasks {
		taskDTO := entity.Task{
			ID:          tasks[i].TaskID,
			UserID:      tasks[i].UserID,
			Title:       tasks[i].Title,
			Description: tasks[i].Description,
			IsCompleted: tasks[i].IsCompleted,
		}
		arr = append(arr, taskDTO)
	}
	return arr
}
