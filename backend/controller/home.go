package controller

import (
	"net/http"
	"social_network/helper"
	"social_network/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	_,_ ,user_id := helper.Auth(DB, r)
	user := models.User{ID: user_id}
	listusers, err := user.GetFollow(DB, 4)

	posts, errpost := user.GetPosts(DB)
	if err != nil && errpost!=nil{
		helper.ErrorPage(w, 500)
		return
	}
	err = helper.WriteJSON(w, http.StatusOK, map[string]interface{}{"success": true, "listesusers":listusers,"posts":posts}, nil)
	if err!=nil {
		helper.ErrorPage(w, 400)
		return
	}
}
