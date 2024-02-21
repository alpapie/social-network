package controller

import (
	"net/http"
	"social_network/helper"
	"social_network/models"
)



func Auth(w http.ResponseWriter, r *http.Request) {
	isAuth, _ ,user_id:= helper.Auth(DB, r)
	if isAuth {
		user:=models.User{ID: user_id}
		err:=user.GetUserById(DB,user_id)
		if err!= nil {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte(`{"error":"no access register ","success":false}`))
			return 
		}
		helper.WriteJSON(w,200,map[string]interface{}{"isauth":true,"success":true,"user":user},nil)
		return
	}
	w.WriteHeader(http.StatusForbidden)
	w.Write([]byte(`{"error":"no access register ","success":false}`))
}
