package handler

import (
	"cls-project/db"
	"cls-project/models"
	"encoding/json"
	"net/http"
)

// @Summary Create a new class
// @Description Create a class user with the provided data
// @Tags class
// @Accept json
// @Produce json
// @param class body models.Class true "Class Data"
// @Success 200 {object} models.Class
// @Router /create-class [post]
func CreateClass(w http.ResponseWriter, r *http.Request) {
	var class models.Class
	json.NewDecoder(r.Body).Decode(&class)

	db := db.GetDB()
	db.Create(&class)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(class)
}

// @Summary Get all classes
// @Description Get a list of classes
// @Tags class
// @Accept json
// @Produce json
// @Success 200 {array} models.Class
// @Router /get-all-classes [get]
func GetAllClasses(w http.ResponseWriter, r *http.Request) {
	db := db.GetDB()
	var classes []models.Class

	db.Find(&classes)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(classes)
}

// @Summary Get class by ID
// @Description Get a specific class by ID
// @Tags class
// @Accept json
// @Produce json
// @Param id query int true "Class ID"
// @Success 200 {object} models.Class
// @Router /get-classes [get]
func GetClassByID(w http.ResponseWriter, r *http.Request) {
	var class models.Class
	id := r.URL.Query().Get("id")

	db := db.GetDB()
	db.Where("id =?", id).First(&class)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(class)
}

// @Summary Get Status for User
// @Description Get Status for all user Class
// @Tags class
// @Accept json
// @Produce json
// @Param id query int true "User ID"
// @Success 200 {array} models.Class
// @Router /user-classes [get]
func GetUserClasses(w http.ResponseWriter, r *http.Request) {
	var class []models.Class
	id := r.URL.Query().Get("id")

	db := db.GetDB()
	db.Preload("Status", "user_id = ?", id).Find(&class)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(class)
}
