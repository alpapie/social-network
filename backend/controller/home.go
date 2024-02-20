package controller

import (
	"fmt"
	"net/http"
	"social_network/helper"
	"social_network/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	_, _, user_id := helper.Auth(DB, r)
	user := models.User{ID: user_id}
	notif := models.Notification{}

	listusers, err := user.GetUnFollow(DB, 4)
	followerAndfollowing, errfol_wing := user.GetFollowerAndFollowing(DB, user_id)
	notifications,errNotif := notif.GetNotf(DB, user_id)

	user.ID=user_id
	posts, errpost := user.GetPosts(DB)

	if err != nil || errpost != nil || errfol_wing != nil || errNotif!=nil {
		fmt.Println(errfol_wing)
		helper.ErrorPage(w, 500)
		return
	}
	err = helper.WriteJSON(w, http.StatusOK, map[string]interface{}{"success": true, "folow_and_following": followerAndfollowing, "listesusers": listusers, "posts": posts,"notifications":notifications}, nil)
	if err != nil {
		fmt.Println(err)
		helper.ErrorPage(w, 400)
		return
	}
}
