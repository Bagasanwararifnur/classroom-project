package handler

import (
	"cls-project/db"
	"cls-project/models"
	"encoding/json"
	"net/http"
)

// @Summary Create a new status
// @Description Create a new status with the provided data
// @Tags status
// @Accept json
// @Produce json
// @Param status body models.Status true "Status Data"
// @Success 200 {object} models.Status
// @Router /create-status [post]
func CreateStatus(w http.ResponseWriter, r *http.Request) {
	var status models.Status
	json.NewDecoder(r.Body).Decode(&status)

	db := db.GetDB()
	db.Create(&status)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}
