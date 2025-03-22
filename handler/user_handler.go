package handler

import (
	"cls-project/db"
	"cls-project/models"
	"encoding/json"
	"net/http"
)

// @Summary Create a new user
// @Description Create a new user with the provided data
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User information"
// @Success 201 {object} models.User
// @Router /users [post]
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	db := db.GetDB()
	db.Create(&user)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

}

// @Summary Get all users
// @Description Get a list of users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} models.User
// @Router /users [get]
func GetUser(w http.ResponseWriter, r *http.Request) {
	db := db.GetDB()
	var users []models.User

	db.Find(&users)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// @Summary Get User By ID
// @Description Get Information of User
// @Tags users
// @Accept json
// @Produce json
// @param id query int true "User ID"
// @Success 200 {object} models.User
// @Router /details-user [get]
func GetUserDetails(w http.ResponseWriter, r *http.Request) {
	var user models.User
	id := r.URL.Query().Get("id")

	db := db.GetDB()
	db.Where("id = ?", id).First(&user)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// @Summary Get All Thread Posted By User
// @Description User All Thread
// @Tags users
// @Accept json
// @Produce json
// @Param id query int true "User ID"
// @Success 200 {object} models.User
// @Router /user-thread [get]
func GetUserThreads(w http.ResponseWriter, r *http.Request) {
	var user models.User
	id := r.URL.Query().Get("id")

	db := db.GetDB()
	db.Where("id =?", id).Preload("Thread").First(&user)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// @Summary Get all comments
// @Description User All Comments
// @Tags users
// @Accept json
// @Produce json
// @Param id query int true "User ID"
// @Success 200 {object} models.User
// @Router /user-comment [get]
func GetUserComments(w http.ResponseWriter, r *http.Request) {
	var user models.User
	id := r.URL.Query().Get("id")

	db := db.GetDB()
	db.Where("id =?", id).Preload("Comment.ThreadDetails").Preload("Comment").First(&user)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
