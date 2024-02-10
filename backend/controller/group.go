package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	helper "social_network/helper"
	"social_network/models"
	"strconv"
)

type NewGroup struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	NotifStatus bool   `json:"notif_type"`
}

func GetAllNotjoinedGroups(w http.ResponseWriter, r *http.Request) {
	var user = models.User{}
	auth, userEmail := helper.Auth(DB, r)
	if !auth {
		fmt.Println("0")
		return
	}
	err := user.GetUserByEmail(DB, userEmail)
	if err != nil {
		fmt.Println("1")
		fmt.Println(err)
		return
	}
	fmt.Println("======------======")
	fmt.Println(user)
	groups, err := models.GetNotJoinedGroupsByUserID(DB, user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var groupss []NewGroup
	for i := 0; i < len(groups); i++ {
		not, errn := models.GetNotificationByUserIDAndType(DB, user.ID, "follow-Group-ID="+strconv.Itoa(groups[i].ID))
		if errn != nil {
			fmt.Println("df")
			fmt.Println(errn)
			return
		}
		fmt.Println("llll")
		fmt.Println(len(not))
		notifStatus := false
		if len(not) == 0 {
			notifStatus = false
		} else {
			notifStatus = true
		}

		var gr = NewGroup{
			ID:          groups[i].ID,
			UserID:      groups[i].UserID,
			Title:       groups[i].Title,
			Description: groups[i].Description,
			NotifStatus: notifStatus,
		}

		groupss = append(groupss, gr)
	}
	fmt.Println("les notififs")
	fmt.Println(groupss)

	groups2, err := models.GetMemberGroupsByUserID(DB, user.ID)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	follower, err := user.GetFollowers(DB)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	followed, err := user.GetFollowed(DB)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(follower)

	responseMap := map[string]interface{}{
		"joined":    groups2,
		"Notjoined": groupss,
		"follower":  follower,
		"followed":  followed,
		"userid":    user.ID,
	}

	jsonData, err := json.Marshal(responseMap)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func CreateInvitationGroup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("here")
	type GroupInvitationData struct {
		GroupID string `json:"groupId"`
		UserID  string `json:"userId"`
		UserIDs []int  `json:"userIds"`
	}

	var groupData GroupInvitationData

	err := json.NewDecoder(r.Body).Decode(&groupData)
	if err != nil {
		fmt.Println("ggg")
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var user = models.User{}
	userID, eru := strconv.Atoi(groupData.UserID)
	groupID, erg := strconv.Atoi(groupData.GroupID)
	if eru != nil || erg != nil {
		fmt.Println("w")
		return
	}
	fmt.Println(userID)

	errr := user.GetUserById(DB, userID)
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

	for i := 0; i < len(groupData.UserIDs); i++ {
		var notification = models.Notification{}
		notification.SenderID = user.ID
		notification.User_id = groupData.UserIDs[i]
		notification.Type = "inviteTO-Group-ID=" + strconv.Itoa(groupID)
		notification.Status = "false"
		ern := notification.CreateNotification(DB)
		if ern != nil {
			fmt.Println(ern)
		}
	}

	fmt.Println("FIN")
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
	notification.Type = "follow-Group-ID=" + strconv.Itoa(group.ID)
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var newGroup models.Group
	fmt.Println(r)
	err := json.NewDecoder(r.Body).Decode(&newGroup)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		fmt.Println("in decode")
		return
	}

	if newGroup.Title == "" || newGroup.Description == "" {
		http.Error(w, "Title and description cannot be empty", http.StatusBadRequest)
		return
	}

	newGroup.UserID = 1
	lastInsertId, err := newGroup.CreateGroup(DB)
	if err != nil {
		http.Error(w, "Failed to create group", http.StatusInternalServerError)
		fmt.Println(err)
		fmt.Println("in user")

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
