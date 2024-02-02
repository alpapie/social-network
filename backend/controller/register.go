package controller

import (
	"database/sql"
	"fmt"
	"html"
	"net/http"
	"regexp"
	helper "social_network/helper"
	"social_network/models"
	"strconv"
	"strings"

	"golang.org/x/crypto/bcrypt"
)



func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user = models.User{}
	
	erf := r.ParseMultipartForm(64)
	if erf != nil {
		// helper.ErrorPage(w, 500)
		return 
	}
	user.Email = html.EscapeString(strings.TrimSpace(r.FormValue("email")))
	password := html.EscapeString(r.FormValue("password"))
	user.FirstName = html.EscapeString(strings.TrimSpace(r.FormValue("firstname")))
	user.LastName = html.EscapeString(strings.TrimSpace(r.FormValue("lastname")))
	user.BirthName = html.EscapeString(strings.TrimSpace(r.FormValue("birthdate")))
	user.Avatar = html.EscapeString(strings.TrimSpace(r.FormValue("avatar")))
	user.NickName = html.EscapeString(strings.TrimSpace(r.FormValue("nickname")))
	user.AboutMe = html.EscapeString(strings.TrimSpace(r.FormValue("aboutme")))
	if !verifLen(password, user.FirstName, user.LastName, user.Email, user.BirthName) {
		helper.ErrorMessage(w, "Enter at least 2 input characters")
		return
	}
	if !isEmailValid(user.Email) {
		// helper.ErrorMessage(w, "bad format of email")
		return
	}
	haspassword, errCrypt := bcrypt.GenerateFromPassword([]byte(password), 5)
	user.TokenLogin = helper.LognToken(user.Email, haspassword)
	user.IsPublic = 0
	err := user.CreateUser(DB)

	if errCrypt != nil || err != nil {
		if strings.HasPrefix(err.Error(), "UNIQUE constraint failed:") {
			fmt.Println(err.Error())
			contraint := "email or username"
			tab:=strings.Split(err.Error(), ":")
			if len(tab)>1 {
				getContrain:=strings.Split(tab[1],".")
				if len(getContrain)>1 {
					contraint=getContrain[1]
				}
			}
			helper.ErrorMessage(w,  contraint+" already exists")
			return
		} else {
			fmt.Println(err)
			// helper.ErrorPage(w, 400)
			return
		}
	}
	helper.WriteJSON(w, 200, map[string]interface{}{"success": true, "data": ""}, nil)
}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

func verifLen(data ...string) bool {
	for _, v := range data {
		if len(v) < 2 {
			return false
		}
	}
	return true
}