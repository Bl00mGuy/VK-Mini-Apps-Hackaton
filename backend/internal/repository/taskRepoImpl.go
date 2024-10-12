package repository

import (
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/dto"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/entity"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/mapper"
	"gorm.io/gorm"
)

type taskRepository struct {
	db *gorm.DB
	mp *mapper.TaskMapper
}

func NewTaskRepository(db *gorm.DB, mp *mapper.TaskMapper) TaskRepository {
	return &taskRepository{
		db: db,
		mp: mp,
	}
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

func (r *taskRepository) FindByID(findTaskDTO *dto.FindTaskDTO) (*dto.TaskDTO, error) {
	var task entity.Task
	if err := r.db.First(&task, findTaskDTO.TaskID).Error; err != nil {
		return nil, err
	}
	return r.mp.ConvertToDTO(&task), nil
}

func (r *taskRepository) FindAll(findAllTasksDTO *dto.FindAllTasksDTO) ([]dto.TaskDTO, error) {
	var tasks []entity.Task
	if err := r.db.Where("user_id = ?", findAllTasksDTO.UserID).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return r.mp.ConvertToSliceDTO(tasks), nil
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
