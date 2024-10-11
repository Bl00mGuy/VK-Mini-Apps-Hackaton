package service

import (
	"encoding/json"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/entity"
	"net/http"
)

func CreateAvatar(w http.ResponseWriter, r *http.Request) {
	// Пример создания аватара на основе данных
	avatar := entity.Avatar{
		Name:      "Иван Иванов",
		ImageURL:  "https://vk.com/images/avatar.jpg",
		Interests: []string{"Киберспорт", "Программирование"},
	}
	json.NewEncoder(w).Encode(avatar)
}
