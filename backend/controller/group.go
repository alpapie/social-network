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
