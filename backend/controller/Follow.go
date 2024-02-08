package controller

import (
	"net/http"
	"social_network/helper"
	"social_network/models"
	"strconv"
)

func Follow(w http.ResponseWriter, r *http.Request) {
	_,_,user_id:= helper.Auth(DB,r)
	limit,errlim:=strconv.Atoi(r.Form.Get("")) 
	if errlim!=nil {
		helper.ErrorPage(w,400)
		return
	}
	user:= models.User{ID: user_id}
	listusers,err:= user.GetFollow(DB,limit)
	if err!=nil {
		helper.ErrorPage(w,500)
		return
	}
	err=helper.WriteJSON(w,http.StatusOK, map[string]interface{}{"success": true, "data": listusers}, nil)
	if err!=nil {
		helper.ErrorPage(w,500)
		return
	}
}
