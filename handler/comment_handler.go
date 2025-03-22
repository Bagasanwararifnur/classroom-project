package handler

import (
	"cls-project/db"
	"cls-project/models"
	"encoding/json"
	"net/http"
)

// @Summary Create a new comment
// @Description Create a Comment
// @Tags comments
// @Accept json
// @Produce json
// @Param comment body models.Comment true "Comment Detail"
// @Success 200 {object} models.Comment
// @Router /comment-create [post]
func CreateComment(w http.ResponseWriter, r *http.Request) {
	var comment models.Comment
	json.NewDecoder(r.Body).Decode(&comment)

	db := db.GetDB()
	db.Create(&comment)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(comment)
}
