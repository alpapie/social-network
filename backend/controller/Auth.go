package controller

import (
	"fmt"
	"net/http"
	"social_network/helper"
)

func Auth(w http.ResponseWriter, r *http.Request) {
	isAuth, _ := helper.Auth(DB, r)
	fmt.Println(isAuth)
	if isAuth {
		w.WriteHeader(200)
		w.Write([]byte(`{"isauth":true,"success":true}`))
		return
	}
	w.WriteHeader(http.StatusForbidden)
	w.Write([]byte(`{"error":"no access register ","success":false}`))
}
