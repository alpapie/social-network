package controller

import (
	"fmt"
	"net/http"
	"social_network/helper"
	"social_network/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	_,_ ,user_id := helper.Auth(DB, r)
	
	user := models.User{ID: user_id}
	user.GetUserById(DB,user_id)
	fmt.Println("THE USER ID BEFORE ", user.ID)
	listusers, err := user.GetUnFollow(DB, 4)
	fmt.Println("THE USER ID AFTER ", user.ID)

	followerAndfollowing,errfol_wing:=user.GetFollowerAndFollowing(DB,user_id)
	fmt.Println("rer" , followerAndfollowing)
	posts, errpost := user.GetPosts(DB)
	if err != nil || errpost!=nil  || errfol_wing!=nil {
		fmt.Println(errfol_wing)
		helper.ErrorPage(w, 500)
		return
	}
	follower, err := user.GetFollowers(DB)
	if err != nil {
		fmt.Println(err)
		helper.ErrorPage(w,500)
		return
	}

	followed, err := user.GetFollowed(DB)
	if err != nil {
		fmt.Println(err)
		helper.ErrorPage(w,500)
		return
	}
	contacts := append(follower, followed...)

	err = helper.WriteJSON(w, http.StatusOK, map[string]interface{}{"success": true,"folow_and_following":follower, "listesusers":listusers,"posts":posts , "contacts" : contacts , "user" :user}, nil)
	if err!=nil {
		fmt.Println(err)
		helper.ErrorPage(w, 400)
		return
	}
}
