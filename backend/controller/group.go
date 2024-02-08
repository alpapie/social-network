package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	helper "social_network/helper"
	"social_network/models"
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
	groups, err := models.GetNotJoinedGroupsByUserID(DB, 1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(groups)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func CreateGroupHandler(w http.ResponseWriter, r *http.Request) {
	// auth, _ := helper.Auth(DB, r)
	// if !auth {
	// 	http.Error(w, "Unauthorized", http.StatusUnauthorized)
	// 	return
	// }
	var newGroup models.Group
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
	
	// replace after by dynamique  id getter
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
        "status": "success",
        "message": "Group created successfully",
        "groupId": lastInsertId,
    })
}