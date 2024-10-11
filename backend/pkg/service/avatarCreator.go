package service

import (
	"encoding/json"
	"net/http"
)

func CreateAvatar(w http.ResponseWriter, r *http.Request) {
	// Пример создания аватара на основе данных
	avatar := Avatar{
		Name:      "Иван Иванов",
		ImageURL:  "https://vk.com/images/avatar.jpg",
		Interests: []string{"Киберспорт", "Программирование"},
	}
	json.NewEncoder(w).Encode(avatar)
}
