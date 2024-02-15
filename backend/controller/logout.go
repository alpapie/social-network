package controller

import (
	"net/http"
	"social_network/helper"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	sess, err := r.Cookie("sessionId")
	errdelete := helper.DeleteSessio(DB, sess.Value)

	if err != nil || errdelete != nil { // sous la contrainte
		helper.ErrorPage(w, 500)
		return
	}

	helper.WriteJSON(w, http.StatusOK, map[string]interface{}{"data": "logout", "success": true}, nil)
}
