package controller

import (
	"database/sql"
	"fmt"
	"net/http"
	"social_network/helper"
	"social_network/models"
	"strings"
	"time"

	"github.com/gofrs/uuid"
)

var DB *sql.DB

func Login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(64)
	fmt.Println(err)
	if err != nil {
		helper.ErrorPage(w, 400)
		return
	}
	r.FormValue("email")
	user := models.User{}

	token := helper.LognToken(r.FormValue("email"), r.FormValue("password"))

	err = user.GetUserByToken(DB, token)
	if err != nil {
		fmt.Println("no roww", err)
		err_resp := helper.ErrorMessage(w, "email or password incorect")
		if err_resp != nil {
			helper.ErrorPage(w, 500)
			return
		}
		return
	}
	if user.Email == "" {
		err_resp := helper.ErrorMessage(w, "email or password incorect")
		if err_resp != nil {
			helper.ErrorPage(w, 500)
			return
		}
		return
	}
	var u1 = uuid.Must(uuid.NewV4())
	sssid := u1.String() + "-" + strings.ReplaceAll(strings.ReplaceAll(time.Now().GoString(), " ", "-"), ",", "_")

	errss := helper.SessionAddOrUpdate(DB, sssid, user.Email, user.ID)
	if errss != nil {
		fmt.Println(errss)
		helper.ErrorPage(w, 500)
		return
	}

	helper.WriteJSON(w, 200, map[string]interface{}{"success": true, "data": sssid, "user": user}, nil)
}
