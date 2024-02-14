package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	helper "social_network/helper"
	"social_network/models"
)

type NewGroup struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	NotifStatus bool      `json:"notif_type"`
	SuggestList []NewUser `json:"suggests"`
}
type NewUser struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Avatar      string `json:"avatar"`
	Isrequested bool   `json:"is_requested"`
}

func CreateFollowGroup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HERE")
	type GroupFollowData struct {
		GroupID int `json:"groupId"`
		UserID  int `json:"userId"`
	}

	var followData GroupFollowData

	err := json.NewDecoder(r.Body).Decode(&followData)
	if err != nil {
		fmt.Println("ggg")
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("userID", followData.UserID)
	fmt.Println("groupID", followData.GroupID)
	var user = models.User{}

	errr := user.GetUserById(DB, followData.UserID)
	if errr != nil {
		fmt.Println("111")
		return
	}
	fmt.Println("THE USER")
	fmt.Println(user)
	sess, err1 := user.HasActiveSession(DB)
	if err1 != nil || !sess {
		fmt.Println(err1)
		fmt.Println("ff")
		return
	}
	fmt.Println(sess)
	group, errg := models.GetGroupByID(DB, followData.GroupID)
	if errg != nil {
		fmt.Println("t")
		return
	}

	var notification = models.Notification{}
	notification.SenderID = user.ID
	notification.User_id = group.UserID
	notification.Type = "follow-Group"
	notification.Group_id = group.ID
	notification.Status = "false"
	ern := notification.CreateNotification(DB)
	if ern != nil {
		fmt.Println(ern)
	}

}

func CreateGroupHandler(w http.ResponseWriter, r *http.Request) {
	// auth, _ := helper.Auth(DB, r)
	// if !auth {
	// 	http.Error(w, "Unauthorized", http.StatusUnauthorized)
	// 	return
	// }
	auth, userEmail := helper.Auth(DB, r)
	if !auth {
		fmt.Println("Not registered")
		return
	}

	var user = models.User{}

	err := user.GetUserByEmail(DB, userEmail)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	var newGroup models.Group
	fmt.Println(r)
	err = json.NewDecoder(r.Body).Decode(&newGroup)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		fmt.Println("in decode")
		return
	}

	if newGroup.Title == "" || newGroup.Description == "" {
		http.Error(w, "Title and description cannot be empty", http.StatusBadRequest)
		return
	}

	newGroup.UserID = user.ID
	lastInsertId, err := newGroup.CreateGroup(DB)
	if err != nil {
		http.Error(w, "Failed to create group", http.StatusInternalServerError)
		fmt.Println(err)
		fmt.Println("in user")

		return
	}
	er := models.JoinGroup(DB, 1, int(lastInsertId))
	if er != nil {
		fmt.Println(er)
		return
	}


	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "success",
		"message": "Group created successfully",
		"groupId": lastInsertId,
	})
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ON GROUP detail")

	helper.WriteJSON(w, 200, map[string]interface{}{"success": true}, nil)
}
