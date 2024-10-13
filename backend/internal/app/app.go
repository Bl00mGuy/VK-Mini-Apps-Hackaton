package app

import (
	"fmt"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/controller"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/mapper"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/repository/repoImpl"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/services/serviceImpl"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

func Start() {
	var db *gorm.DB
	var retries = 5
	var attempt int
	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		"localhost",
		"db",
		"db",
		"postgres",
		"5432")

	for attempt = 1; attempt <= retries; attempt++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
		if err != nil {
			time.Sleep(5 * time.Second)
			continue
		}
		break
	}

	if attempt > retries {
		log.Fatalf("Failed to connect to database after %d attempts", retries)
		return
	}

	taskMapper := mapper.NewTaskMapper()
	taskRepository := repoImpl.NewTaskRepository(db, taskMapper)
	taskService := serviceImpl.NewTaskService(taskRepository)
	taskHandler := controller.NewTaskHandler(*taskService)

	router := gin.Default()
	taskRoutes := router.Group("/tasks")
	{
		taskRoutes.POST("", taskHandler.CreateTask)
		taskRoutes.GET("/:task_id", taskHandler.FindTaskByID)
		taskRoutes.GET("/all/:user_id", taskHandler.FindAllTasks)
		taskRoutes.PUT("", taskHandler.UpdateTask)
		taskRoutes.DELETE("/:task_id", taskHandler.DeleteTask)
	}

	err = router.Run(":8080")
	if err != nil {
		log.Fatalf("Error in gin run function: %v", err)
		return
	}

	log.Println("Gin server has started")
}
