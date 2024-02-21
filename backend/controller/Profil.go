package controller

import (
	"fmt"
	"net/http"
	"social_network/helper"
	"social_network/models"
	"strconv"
)

func Profil(w http.ResponseWriter, r *http.Request) {
	_, _, user_id := helper.Auth(DB, r)
	user := models.User{ID: user_id}
	err := user.GetUserById(DB, user.ID)
	currentUser := user

	follower, errFollow := user.Folower(DB, user_id)
	following, errFollowing := user.Following(DB, user_id)
	createdPost, errPost := user.CreatedPost(DB, user_id)
	fmt.Println("dreated post",len(createdPost))
	if err != nil || errFollow != nil || errFollowing != nil || errPost != nil {
		fmt.Println(err, errFollow, errFollowing, errPost)
		helper.ErrorPage(w, 500)
		return
	}

	err = helper.WriteJSON(w, http.StatusOK, map[string]interface{}{"success": true, "user": currentUser, "follower": follower, "following": following, "createdPost": createdPost}, nil)
	if err != nil {
		helper.ErrorPage(w, 400)
		return
	}
}

func UpdateProfil(w http.ResponseWriter, r *http.Request) {
	isPublic, err := strconv.Atoi(r.URL.Query().Get("ispublic"))
	if err != nil {
		helper.ErrorPage(w, 400)
		return
	}
	_, _, user_id := helper.Auth(DB, r)
	user := models.User{ID: user_id,IsPublic: isPublic}
	err=user.UpdateStatus(DB)
	if err!=nil {
		helper.ErrorPage(w, 400)
		return
	}

	err = helper.WriteJSON(w, http.StatusOK, map[string]interface{}{"success": true, "user":user_id}, nil)
	if err != nil {
		helper.ErrorPage(w, 400)
		return
	}
}

func FollowerProfil(w http.ResponseWriter, r *http.Request){
	user_id,err:=strconv.Atoi(r.URL.Query().Get("user_id"))
	user := models.User{ID: user_id}
	if err!=nil || user_id<=0{
		fmt.Println("user id profil follow error",user_id)
		helper.ErrorPage(w, 400)
		return
	}
	err = user.GetUserById(DB, user.ID)
	currentUser := user

	follower, errFollow := user.Folower(DB, user_id)
	following, errFollowing := user.Following(DB, user_id)
	createdPost, errPost := user.CreatedPost(DB, user_id)

	if err != nil || errFollow != nil || errFollowing != nil || errPost != nil {
		fmt.Println(err, errFollow, errFollowing, errPost)
		helper.ErrorPage(w, 500)
		return
	}

	err = helper.WriteJSON(w, http.StatusOK, map[string]interface{}{"success": true, "user": currentUser, "follower": follower, "following": following, "createdPost": createdPost}, nil)
	if err != nil {
		helper.ErrorPage(w, 400)
		return
	}
}
