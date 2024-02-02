package controller

import (
	"database/sql"
	"fmt"
	"net/http"
	"social_network/helper"
	"social_network/models"
)

var DB *sql.DB

func Login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(64)
	fmt.Println(err)
	if err != nil {
		helper.ErrorPage(w, 400)
	}
	r.FormValue("email")
	user := models.User{}
	// user.
	fmt.Println(r.FormValue("email"))
	fmt.Println(r.FormValue("password"))
	token := helper.LognToken(r.FormValue("email"), r.FormValue("password"))
	fmt.Println(token)
	err=user.GetUserByToken(DB, token)
	if err != nil {
		helper.ErrorPage(w, 400)
	}
}
