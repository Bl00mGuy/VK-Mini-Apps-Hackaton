package main

import (
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/pkg/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	// TODO: FIX HANDLERS!!!
	r.HandleFunc("/api/avatar", service.CreateAvatar()).Methods("POST")
	r.HandleFunc("/api/achievements", GetAchievements).Methods("GET")

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
