package app

import (
	"errors"
	"fmt"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/controller"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/factory"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

const (
	domainName    = "localhost"
	userForDB     = "db"
	passwordForDB = "db"
	dbName        = "postgres"
	portForDB     = "5432"
	portForServer = ":80"
)

func Start() {
	var db *gorm.DB
	var retries = 5
	var attempt int
	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		domainName,
		userForDB,
		passwordForDB,
		dbName,
		portForDB)

	err = tryToConnectToDB(attempt, retries, db, err, dsn)
	if err != nil {
		return
	}

	router := gin.Default()
	getAllRoutes(router, db)

	err = router.Run(portForServer)
	if err != nil {
		log.Fatalf("Error in gin run function: %v", err)
		return
	}

	log.Println("Gin server has started")
}

func tryToConnectToDB(attempt, retries int, db *gorm.DB, err error, dsn string) error {
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
		return errors.New("failed to connect to database")
	}

	return nil
}

func getAllRoutes(router *gin.Engine, db *gorm.DB) {
	taskHandler := factory.CreateTaskHandler(db)
	pingHandler := factory.CreatePingHandler()

	getTaskRoutes(router, taskHandler)
	getPingRoutes(router, pingHandler)
}

func getTaskRoutes(router *gin.Engine, taskHandler *controller.TaskHandler) {
	taskRoutes := router.Group("/tasks")
	{
		taskRoutes.POST("", taskHandler.CreateTask)
		taskRoutes.GET("/:task_id", taskHandler.FindTaskByID)
		taskRoutes.GET("/all/:user_id", taskHandler.FindAllTasks)
		taskRoutes.PUT("", taskHandler.UpdateTask)
		taskRoutes.DELETE("/:task_id", taskHandler.DeleteTask)
	}
}

func getPingRoutes(router *gin.Engine, pingHandler *controller.PingHandler) {
	pingRoutes := router.Group("/ping")
	{
		pingRoutes.GET("", pingHandler.Ping)
	}
}
