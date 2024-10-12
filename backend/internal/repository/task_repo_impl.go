package repository

import (
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/dto"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/entity"
	"gorm.io/gorm"
)

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db}
}

func (r *taskRepository) Create(createTaskDTO *dto.CreateTaskDTO) error {
	task := &entity.Task{
		UserID:      createTaskDTO.UserID,
		Title:       createTaskDTO.Title,
		Description: createTaskDTO.Description,
		IsCompleted: false,
	}
	return r.db.Create(task).Error
}

func (r *taskRepository) FindByID(findTaskDTO *dto.FindTaskDTO) (*entity.Task, error) {
	var task entity.Task
	if err := r.db.First(&task, findTaskDTO.TaskID).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *taskRepository) FindAll(findAllTasksDTO *dto.FindAllTasksDTO) ([]entity.Task, error) {
	var tasks []entity.Task
	if err := r.db.Where("user_id = ?", findAllTasksDTO.UserID).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *taskRepository) Update(updateTaskDTO *dto.UpdateTaskDTO) error {
	var task entity.Task
	if err := r.db.First(&task, updateTaskDTO.TaskID).Error; err != nil {
		return err
	}

	task.Title = updateTaskDTO.Title
	task.Description = updateTaskDTO.Description
	task.IsCompleted = updateTaskDTO.IsCompleted

	return r.db.Save(task).Error
}

func (r *taskRepository) Delete(deleteTaskDTO *dto.DeleteTaskDTO) error {
	return r.db.Delete(&entity.Task{}, deleteTaskDTO.TaskID).Error
}
