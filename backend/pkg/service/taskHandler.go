package service

import (
	"encoding/json"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/entity"
	"net/http"
)

func GenerateTask(w http.ResponseWriter, r *http.Request) {
	task := entity.Task{
		Title:       "Участие в вебинаре",
		Description: "Примите участие в вебинаре и получите 100 очков",
		Points:      100,
	}
	json.NewEncoder(w).Encode(task)
}
