package controller

import (
	"database/sql"
	"fmt"
	"net/http"
	"social_network/helper"
	"social_network/models"
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
	// user.
	fmt.Println(r.FormValue("email"))
	fmt.Println(r.FormValue("password"))
	token := helper.LognToken(r.FormValue("email"), r.FormValue("password"))
	fmt.Println(token)
	err = user.GetUserByToken(DB, token)
	if err != nil {
		fmt.Println("no roww",err)
		helper.ErrorPage(w, 400)
		return
	}
	if(user.Email==""){
		err_resp:= helper.ErrorMessage(w,"email or password incorect")
		if err_resp!=nil {
			helper.ErrorPage(w,500)
			return
		}
		return
	}
	var u1 = uuid.Must(uuid.NewV4())
	sssid := u1.String() + "-" + time.Now().GoString()
	cookie := http.Cookie{
		Name:     "sessionid",
		Value:    sssid,
		Expires:  time.Now().Add(time.Hour * 24 * 3),
		Path:     "/",
		MaxAge:   3600 * 24 * 3,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &cookie)

	errss := helper.SessionAddOrUpdate(DB, sssid, user.Email)
	if errss != nil {
		fmt.Println(errss)
		helper.ErrorPage(w, 500)
		return
	}

	helper.WriteJSON(w, 200, map[string]interface{}{"success": true, "data": cookie}, nil)
}
