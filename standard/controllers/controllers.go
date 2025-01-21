package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/HarshThakur1509/boilerplate/standard/initializers"
)

type HealthResponse struct {
	Status string `json:"status"`
	DB     string `json:"database"`
}

func Health(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{Status: "UP", DB: "OK"}

	// Ping the database
	sqlDB, err := initializers.DB.DB()
	if err != nil || sqlDB.Ping() != nil {
		response.DB = "DOWN"
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	// Write the JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
