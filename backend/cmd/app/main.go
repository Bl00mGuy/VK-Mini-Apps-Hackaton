package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	// Пример маршрутов
	r.HandleFunc("/api/avatar", CreateAvatar).Methods("POST")
	r.HandleFunc("/api/achievements", GetAchievements).Methods("GET")

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
