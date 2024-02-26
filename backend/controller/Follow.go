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

	notificationType := "follow"

	if follower.IsPublic == 1 {
		errreq := follower.AddFollower(DB, user_id)
		if errreq != nil {
			if errreq.Error() == "already followed" {
				fmt.Println(errreq)
				helper.ErrorPage(w, 400)
				return
			}
			helper.ErrorPage(w, 500)
			return
		}
		notificationType="succesrequest"
	} 
	
	newNotification := models.Notification{User_id: follower.ID, SenderID: current_user.ID, FirstName: current_user.FirstName, Status: "0", LastName: current_user.LastName, Avatar: current_user.Avatar, Type: notificationType}
	errNotification := newNotification.CreateNotification(DB)
	if errNotification != nil {
		helper.ErrorPage(w, 500)
		return
	}
	SendSocketNotification(newNotification, "notification")

	err = helper.WriteJSON(w, http.StatusOK, map[string]interface{}{"success": true, "user_id": user_id}, nil)
	if err != nil {
		helper.ErrorPage(w, 400)
		return
	}
}

func NotificationRequestTraiment(w http.ResponseWriter, r *http.Request) {
	follow_id, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	notificationId, errnotifid := strconv.Atoi(r.URL.Query().Get("notif_id"))
	accept, erraccept := strconv.Atoi(r.URL.Query().Get("accept"))
	group_id, errgroup_id := strconv.Atoi(r.URL.Query().Get("group_id"))

	if err != nil || errnotifid != nil || erraccept != nil || errgroup_id != nil {
		fmt.Println(err, errnotifid, erraccept, errgroup_id)
		helper.ErrorPage(w, 400)
		return
	}

	_, _, user_id := helper.Auth(DB, r)

	newNotification := models.Notification{}
	errGettingNotifs := newNotification.GetNotificationByUserIDAndTypeAndnotifid(DB, notificationId, follow_id, user_id, group_id)

	if errGettingNotifs != nil || newNotification.ID == 0 {
		fmt.Println("error get notif", errGettingNotifs)
		helper.ErrorPage(w, 500)
		return
	}

	if accept == 1 {
		errTra := TraiteNotif(user_id, follow_id, group_id)
		if errTra != nil {
			if errTra.Error() == "already followed" {
				fmt.Println(errTra)
				helper.ErrorPage(w, 400)
				return
			}
			fmt.Println("error traitement", errTra)
			helper.ErrorPage(w, 500)
			return
		}

		fmt.Println(newNotification)
		newNotification.User_id = newNotification.SenderID
		newNotification.SenderID = user_id
		newNotification.Type = "succesrequest"
		err = newNotification.CreateNotification(DB)
		if err != nil {
			fmt.Println("error create notif", err)
			helper.ErrorPage(w, 500)
			return
		}
	}

	err = newNotification.MarkAsRead(DB)
	if err != nil {
		fmt.Println("error create mask notf", err)
		helper.ErrorPage(w, 500)
		return
	}

	SendSocketNotification(newNotification, "notification")
	err = helper.WriteJSON(w, http.StatusOK, map[string]interface{}{"success": true, "user_id": user_id}, nil)
	if err != nil {
		helper.ErrorPage(w, 400)
		return
	}
}

func TraiteNotif(user_id, follow_id, group_id int) error {
	if group_id > 0 {
		return models.JoinGroup(DB, follow_id, group_id)
	}

	current_user := models.User{ID: user_id}
	return current_user.AddFollower(DB, follow_id)

}

func MarkNotificationAsRead(w http.ResponseWriter, r *http.Request) {
	fmt.Println("\nmarking as read\n")
	follow_id, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		helper.ErrorPage(w, 400)
		return
	}

	notificationId, err := strconv.Atoi(r.URL.Query().Get("notif_id"))
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

	newNotification := models.Notification{ID: notificationId, User_id: current_user.ID, SenderID: follower.ID, FirstName: current_user.FirstName, LastName: current_user.LastName, Avatar: current_user.Avatar, Type: ""}
	newNotification.MarkAsRead(DB)
	if err != nil {
		helper.ErrorPage(w, 500)
		return
	}

	err = helper.WriteJSON(w, http.StatusOK, map[string]interface{}{"success": true, "user_id": user_id}, nil)
	if err != nil {
		helper.ErrorPage(w, 400)
		return
	}
}
