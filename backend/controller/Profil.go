package controller

import (
	"net/http"
	"social_network/helper"
	"social_network/models"
)

func Profil(w http.ResponseWriter, r *http.Request) {
	_, _, user_id := helper.Auth(DB, r)
	user := models.User{ID: user_id}
	err:=user.GetUserById(DB,user_id)
	if err!=nil {
		if err != nil{
			helper.ErrorPage(w, 500)
			return
		}
	}
	err = helper.WriteJSON(w, http.StatusOK, map[string]interface{}{"success": true, "user":user,"follow":nil,"following":nil}, nil)
	if err!=nil {
		helper.ErrorPage(w, 400)
		return
	}
}

func UpdateProfil(w http.ResponseWriter, r *http.Request) {
	
}
