package dto

type CreateTaskDTO struct {
	UserID      uint   `json:"user_id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type FindTaskDTO struct {
	TaskID uint `json:"task_id" binding:"required"`
}

type FindAllTasksDTO struct {
	UserID uint `json:"user_id" binding:"required"`
}

type TaskDTO struct {
	TaskID      uint   `json:"task_id" binding:"required"`
	UserID      uint   `json:"user_id" binding:"required"`
	Title       string `json:"title"   binding:"required"`
	Description string `json:"description" binding:"required"`
	IsCompleted bool   `json:"is_completed" binding:"required"`
}

type UpdateTaskDTO struct {
	TaskID      uint   `json:"task_id" binding:"required"`
	UserID      uint   `json:"user_id" binding:"required"`
	Title       string `json:"title"   binding:"required"`
	Description string `json:"description" binding:"required"`
	IsCompleted bool   `json:"is_completed" binding:"required"`
}

type DeleteTaskDTO struct {
	TaskID uint `json:"task_id" binding:"required"`
}
