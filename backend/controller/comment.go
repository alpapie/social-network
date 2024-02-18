package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	helper "social_network/helper"
	"social_network/models"
)

func AddCommentHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, _, userId := helper.Auth(DB, r)

	newComment := models.Comment{
		UserId: userId,
	}
	err := json.NewDecoder(r.Body).Decode(&newComment)
	fmt.Println("here is the comment ", newComment)
	if err != nil || !newComment.ValidateComment() {
		helper.ErrorPage(w, 400)
		return
	}

	res, err := newComment.IsAllowedToComment(DB)
	if err != nil {
		helper.ErrorPage(w, 500)
		return
	}

	if res == 0 {
		helper.ErrorPage(w, 401)
		return
	}

	newCommentId, err := newComment.Create(DB)
	if err != nil {
		helper.ErrorPage(w, 500)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{"success": newCommentId})
}
