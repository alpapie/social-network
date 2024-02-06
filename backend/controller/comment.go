package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"social_network/helper"
	"social_network/models"
)

func AddCommentHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	newComment := models.Comment{
		UserId: 2,
	}
	err := json.NewDecoder(r.Body).Decode(&newComment)
	fmt.Println("here is the comment ", newComment)
	if err != nil || !newComment.ValidateComment() {
		helper.Error(w, err, 400)
		return
	}

	res, err := newComment.IsAllowedToComment(DB)
	if err != nil {
		helper.Error(w, err, 500)
		return
	}

	if res == 0 {
		helper.Error(w, err, 401)
		return
	}

	newCommentId, err := newComment.Create(DB)
	if err != nil {
		helper.Error(w, err, 500)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{"success": newCommentId})
}
