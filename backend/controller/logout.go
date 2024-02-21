package controller

import (
	"fmt"
	"net/http"
	"social_network/helper"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	sess, err := r.Cookie("sessionId")
	if err !=nil {
		fmt.Println(err)
		helper.ErrorPage(w, 500)
		return
	}
	errdelete := helper.DeleteSessio(DB, sess.Value)

	if  errdelete != nil { // sous la contrainte
		helper.ErrorPage(w, 500)
		return
	}

	helper.WriteJSON(w, http.StatusOK, map[string]interface{}{"data": "logout", "success": true}, nil)
}
