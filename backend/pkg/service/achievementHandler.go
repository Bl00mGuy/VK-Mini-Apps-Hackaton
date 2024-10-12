package service

import (
	"encoding/json"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/entity"
	"net/http"
)

func GetAchievements(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	var achievements []entity.Achievement
	db.Where("user_id = ?", userID).Find(&achievements)
	json.NewEncoder(w).Encode(achievements)
}
