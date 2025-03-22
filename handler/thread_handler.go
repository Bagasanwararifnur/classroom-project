package handler

import (
	"cls-project/db"
	"cls-project/models"
	"encoding/json"
	"net/http"
)

// @Summary Create a new thread
// @Description Create a new thread with the provided data
// @Tags threads
// @Accept json
// @Produce json
// @Param thread body models.Thread true "User information"
// @Success 201 {object} models.Thread
// @Router /thread-create [post]
func CreateThread(w http.ResponseWriter, r *http.Request) {
	var thread models.Thread
	json.NewDecoder(r.Body).Decode(&thread)

	db := db.GetDB()
	db.Create(&thread)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(thread)
}

// @Summary Get all threads
// @Description Get a list of threads
// @Tags threads
// @Accept json
// @Produce json
// @Success 201 {array} models.Thread
// @Router /thread-get [get]
func GetThreads(w http.ResponseWriter, r *http.Request) {
	db := db.GetDB()
	var threads []models.Thread

	db.Find(&threads)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(threads)
}

// @Summary Get Thread By ID
// @Description Get Information of Thread
// @Tags threads
// @Accept json
// @Produce json
// @Param id query int true "Thread ID"
// @Success 200 {object} models.Thread
// @Router /get-thread [get]
func GetThreadByID(w http.ResponseWriter, r *http.Request) {
	var thread models.Thread
	id := r.URL.Query().Get("id")

	db := db.GetDB()
	db.Where("id =?", id).First(&thread)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(thread)
}

// @Summary Get Thread with All Comment
// @Description Get Thread Page
// @Tags threads
// @Accept json
// @Produce json
// @param id query int true "Thread ID"
// @Success 200 {Object} models.Thread
// @Router /thread-page [get]
func GetThreadPage(w http.ResponseWriter, r *http.Request) {
	var thread models.Thread
	id := r.URL.Query().Get("id")

	db := db.GetDB()
	db.Where("id =?", id).Preload("Comment").Find(&thread)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(thread)
}
