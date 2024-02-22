package controller

import (
	"fmt"
	"net/http"
	"social_network/helper"
	"social_network/models"
	"strconv"
)

func UnFollow(w http.ResponseWriter, r *http.Request) {

	_, _, user_id := helper.Auth(DB, r)

	user := models.User{ID: user_id}
	listusers, err := user.GetUnFollow(DB, 100)
	if err != nil {
		helper.ErrorPage(w, 500)
		return
	}
	err = helper.WriteJSON(w, http.StatusOK, map[string]interface{}{"success": true, "listusers": listusers}, nil)
	if err != nil {
		helper.ErrorPage(w, 500)
		return
	}
}

func Follow(w http.ResponseWriter, r *http.Request) {

	follow_id, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		helper.ErrorPage(w, 400)
		return
	}
	_, _, user_id := helper.Auth(DB, r)

	follower := models.User{}
	current_user := models.User{}
	errfollow := follower.GetUserById(DB, follow_id)
	erruser := current_user.GetUserById(DB, user_id)

	if errfollow != nil || erruser != nil {
		helper.ErrorPage(w, 400)
		return
	}

	notificationType := ""

	if follower.IsPublic == 1 {
		errreq := follower.AddFollower(DB, user_id)
		if errreq != nil {
			helper.ErrorPage(w, 500)
			return
		}
		notificationType = "follow"
	} else {
		notificationType = "followRequest"
		fmt.Println("sending notification for private profile")
	}

	newNotification := models.Notification{User_id: follower.ID, SenderID: current_user.ID, FirstName: follower.FirstName, Status: "0", LastName: follower.LastName, Avatar: follower.Avatar, Type: notificationType}
	errNotification := newNotification.CreateNotification(DB)
	if errNotification != nil {
		helper.ErrorPage(w, 500)
		return
	}
	SendSocketMessage(w, follow_id, newNotification, "notification")

	err = helper.WriteJSON(w, http.StatusOK, map[string]interface{}{"success": true, "user_id": user_id}, nil)
	if err != nil {
		helper.ErrorPage(w, 400)
		return
	}
}

func AcceptFollowRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Follow request accepted!")
	follow_id, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		helper.ErrorPage(w, 400)
		return
	}

	_, _, user_id := helper.Auth(DB, r)

	follower := models.User{}
	current_user := models.User{}
	errFollow := follower.GetUserById(DB, follow_id)
	errUser := current_user.GetUserById(DB, user_id)

	if errFollow != nil || errUser != nil {
		helper.ErrorPage(w, 400)
		return
	}

	errReq := follower.AddFollower(DB, user_id)
	if errReq != nil {
		helper.ErrorPage(w, 500)
		return
	}

	currentNotificationsByType, errGettingNotifs := models.GetNotificationByUserIDAndType(DB, follow_id, user_id, "followRequest", 0)
	if errGettingNotifs != nil {
		helper.ErrorPage(w, 500)
		return
	}

	if len(currentNotificationsByType) > 0 {
		err := currentNotificationsByType[0].MarkAsRead(DB, user_id)
		if err != nil {
			helper.ErrorPage(w, 500)
			return
		}
	}

	newNotification := models.Notification{User_id: current_user.ID, SenderID: follower.ID, FirstName: follower.FirstName, LastName: follower.LastName, Avatar: follower.Avatar, Type: "acceptedFollowRequest"}
	SendSocketMessage(w, follow_id, newNotification, "notification")
}

func DeclineFollowRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Follow request declined!")
	follow_id, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		helper.ErrorPage(w, 400)
		return
	}
	_, _, user_id := helper.Auth(DB, r)

	follower := models.User{}
	current_user := models.User{}
	errFollow := follower.GetUserById(DB, follow_id)
	errUser := current_user.GetUserById(DB, user_id)

	if errFollow != nil || errUser != nil {
		helper.ErrorPage(w, 400)
		return
	}

	currentNotificationsByType, errGettingNotifs := models.GetNotificationByUserIDAndType(DB, follow_id, user_id, "followRequest", 0)
	if errGettingNotifs != nil {
		helper.ErrorPage(w, 500)
		return
	}

	if len(currentNotificationsByType) > 0 {
		err := currentNotificationsByType[0].MarkAsRead(DB, user_id)
		if err != nil {
			helper.ErrorPage(w, 500)
			return
		}
	}

	fmt.Println("----------------------------User Notifications---------------------------")
	fmt.Println("length", len(currentNotificationsByType), "content", currentNotificationsByType[0])
	fmt.Println("--------------------------------------------------------------------------")

}
