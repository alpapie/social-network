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

	// followerAndfollowing, errfol_wing := user.GetFollowerAndFollowing(DB, user_id)
	notifications, errNotif := notif.GetNotification(DB, user_id)

	user.GetUserById(DB,user_id)
	listusers, err := user.GetUnFollow(DB, 4)
	follower, errfolow := user.GetFollowers(DB)
	followed, errfolowed := user.GetFollowed(DB)

	user.ID = user_id
	posts, errpost := user.GetPosts(DB)

	if err != nil || errpost != nil || errfolowed != nil || errNotif != nil || errfolow!=nil{
	fmt.Println(err  ,errpost  ,errfolowed  ,errNotif  ,errfolow)
		helper.ErrorPage(w, 500)
		return
	}
	contacts := append(follower, followed...)

	err = helper.WriteJSON(w, http.StatusOK, map[string]interface{}{"success": true,"folow_and_following":follower, "listesusers":listusers,"posts":posts ,"notifications":notifications, "contacts" : contacts , "user" :user}, nil)
	if err!=nil {
		fmt.Println(err)
		helper.ErrorPage(w, 400)
		return
	}
}
