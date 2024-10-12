package controller

import (
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/dto"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/services/impl"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TaskHandler struct {
	taskService impl.TaskService
}

func NewTaskHandler(taskService impl.TaskService) *TaskHandler {
	return &TaskHandler{
		taskService: taskService,
	}
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var createTaskDTO dto.CreateTaskDTO
	if err := c.ShouldBindJSON(&createTaskDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.taskService.CreateTask(&createTaskDTO); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Task created successfully", "task": createTaskDTO})
}

func (h *TaskHandler) FindTaskByID(c *gin.Context) {
	var findTaskDTO dto.FindTaskDTO
	if err := c.ShouldBindUri(&findTaskDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	taskDTO, err := h.taskService.FindTaskByID(&findTaskDTO)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, taskDTO)
}

func (h *TaskHandler) FindAllTasks(c *gin.Context) {
	var findAllTasksDTO dto.FindAllTasksDTO
	if err := c.ShouldBindJSON(&findAllTasksDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tasks, err := h.taskService.FindAllTasks(&findAllTasksDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tasks"})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) UpdateTask(c *gin.Context) {
	var updateTaskDTO dto.UpdateTaskDTO
	if err := c.ShouldBindJSON(&updateTaskDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.taskService.UpdateTask(&updateTaskDTO); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully", "task": updateTaskDTO})
}

func (h *TaskHandler) DeleteTask(c *gin.Context) {
	var deleteTaskDTO dto.DeleteTaskDTO
	if err := c.ShouldBindUri(&deleteTaskDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.taskService.DeleteTask(&deleteTaskDTO); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
