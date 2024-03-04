package controller

import (
	"fmt"
	"net/http"
	"social_network/helper"
	"social_network/models"
	"strconv"
)

func UnFollowUserHandler(w http.ResponseWriter, r *http.Request) {
	userToUnfollowId, err := strconv.Atoi(r.URL.Query().Get("userToUnfollow_id"))
	if err != nil {
		fmt.Println("1")
		helper.ErrorPage(w, 400)
		return
	}

	_, _, currentUserId := helper.Auth(DB, r)
	currentUser := models.User{}

	errGettingUser := currentUser.GetUserById(DB, currentUserId)
	if errGettingUser != nil {
		fmt.Println("2")
		helper.ErrorPage(w, 400)
		return
	}

	errUnfollowing := currentUser.UnFollowUser(DB, userToUnfollowId)
	if errUnfollowing != nil {
		if errUnfollowing.Error() == "not following" {
			fmt.Println("3")
			helper.ErrorPage(w, 400)
			return
		}
		fmt.Println("4")
		helper.ErrorPage(w, 500)
		return
	}

	if errDeleteNotif := models.DeleteNotificationByUsers(DB, userToUnfollowId, currentUserId); errDeleteNotif != nil {
		fmt.Println("error deleting notif by users", err)
		helper.ErrorPage(w, 500)
		return
	}

	err = helper.WriteJSON(w, http.StatusOK, map[string]interface{}{"success": true, "user_id": userToUnfollowId}, nil)
	if err != nil {
		fmt.Println("5")
		helper.ErrorPage(w, 500)
		return
	}
}
