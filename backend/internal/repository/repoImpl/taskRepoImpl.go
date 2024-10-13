package repoImpl

import (
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/dto"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/entity"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/mapper"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/repository"
	"gorm.io/gorm"
)

type taskRepository struct {
	db *gorm.DB
	mp *mapper.TaskMapper
}

func NewTaskRepository(db *gorm.DB, mp *mapper.TaskMapper) repository.TaskRepository {
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

func (r *taskRepository) FindByID(id uint) (*dto.TaskDTO, error) {
	var task entity.Task
	if err := r.db.First(&task, id).Error; err != nil {
		return nil, err
	}
	return r.mp.ConvertToDTO(&task), nil
}

func (r *taskRepository) FindAll(userID uint) ([]dto.TaskDTO, error) {
	var tasks []entity.Task
	if err := r.db.Where("user_id = ?", userID).Find(&tasks).Error; err != nil {
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

func (r *taskRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Task{}, id).Error
}
