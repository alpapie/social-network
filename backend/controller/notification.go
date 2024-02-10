package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	helper "social_network/helper"
	"social_network/models"
	"strconv"
)

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
		not, errn := models.GetNotificationByUserIDAndType(DB, user.ID, groups[i].UserID, "follow-Group-ID="+strconv.Itoa(groups[i].ID))
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
	friends := append(follower, followed...)

	groups2, err := models.GetMemberGroupsByUserID(DB, user.ID)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var groupss2 []NewGroup
	for i := 0; i < len(groups2); i++ {
		var friendss []NewUser
		for j := 0; j < len(friends); j++ {
			notfollow, errn := models.GetNotificationByUserIDAndType(DB, user.ID, friends[j].ID, "inviteTO-Group-ID="+strconv.Itoa(groups2[i].ID))
			if errn != nil {
				fmt.Println("df")
				fmt.Println(errn)
				return
			}
			requested := false
			if len(notfollow) == 0 {
				requested = false
			} else {
				requested = true
			}
			var fr = NewUser{
				ID:          friends[j].ID,
				FirstName:   friends[j].FirstName,
				LastName:    friends[j].LastName,
				Email:       friends[j].Email,
				Avatar:      friends[j].Avatar,
				Isrequested: requested,
			}
			friendss = append(friendss, fr)
		}
		var group = NewGroup{
			ID:          groups2[i].ID,
			UserID:      groups2[i].UserID,
			Title:       groups2[i].Title,
			Description: groups2[i].Description,
			SuggestList: friendss,
		}
		groupss2 = append(groupss2, group)

	}

	fmt.Println("-----------------")
	fmt.Println(groupss2)
	fmt.Println("-----------------")

	responseMap := map[string]interface{}{
		"joined":    groupss2,
		"Notjoined": groupss,
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
