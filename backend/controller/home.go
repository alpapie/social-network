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
	fmt.Println("THE USER ID BEFORE ", user.ID)
	listusers, err := user.GetUnFollow(DB, 4)
	fmt.Println("THE USER ID AFTER ", user.ID)

	followerAndfollowing,errfol_wing:=user.GetFollowerAndFollowing(DB,user_id)

	posts, errpost := user.GetPosts(DB)
	if err != nil || errpost!=nil || errfol_wing!=nil{
		fmt.Println(errfol_wing)
		helper.ErrorPage(w, 500)
		return
	}
	err = helper.WriteJSON(w, http.StatusOK, map[string]interface{}{"success": true,"folow_and_following":followerAndfollowing, "listesusers":listusers,"posts":posts}, nil)
	if err!=nil {
		fmt.Println(err)
		helper.ErrorPage(w, 400)
		return
	}
}
