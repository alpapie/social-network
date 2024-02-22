package controller

import (
	"fmt"
	"net/http"
	helper "social_network/helper"
	"social_network/models"
	"strconv"
)

func GetGroupMessageHandler(w http.ResponseWriter, r *http.Request) {
	groupId, er := strconv.Atoi(r.URL.Query().Get("groupid"))
	if er != nil {
		helper.ErrorPage(w, 500)
		return
	}
	_, _, UserID := helper.Auth(DB, r)
	user := models.User{ID: UserID}
	checker, errc := user.IsGroupmemeber(DB, groupId)
	if errc != nil || !checker {
		helper.ErrorPage(w, 400)
		return
	}
	messages, errm := models.GetGroupDiscussion(DB, groupId)
	if errm != nil {
		helper.ErrorPage(w, 500)
		return
	}
	err := helper.WriteJSON(w, http.StatusOK, map[string]interface{}{"success": true, "messages": messages}, nil)
	if err != nil {
		fmt.Println(err)
		helper.ErrorPage(w, 400)
		return
	}
}
func GetChatMessageHandler(w http.ResponseWriter, r *http.Request) {
	receiverId, er := strconv.Atoi(r.URL.Query().Get("receiver"))
	if er != nil {
		helper.ErrorPage(w, 500)
		return
	}
	_, _, UserID := helper.Auth(DB, r)
	checker, errc := models.AreFriend(DB, receiverId, UserID)
	if errc != nil || !checker {
		helper.ErrorPage(w, 500)
		return
	}

	messages, errm := models.GetDiscussion(DB, receiverId, UserID)
	if errm != nil {
		helper.ErrorPage(w, 500)
		return
	}
	err := helper.WriteJSON(w, http.StatusOK, map[string]interface{}{"success": true, "messages": messages}, nil)
	if err != nil {
		fmt.Println(err)
		helper.ErrorPage(w, 400)
		return
	}
}