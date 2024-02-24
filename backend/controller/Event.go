package controller

import (
	"fmt"
	"html"
	"net/http"
	helper "social_network/helper"
	"social_network/models"
	"strconv"
	"strings"
)


func CreateEvent(w http.ResponseWriter, r *http.Request) {
	var user = models.User{}
	_, userEmail, _ := helper.Auth(DB, r)

	err := user.GetUserByEmail(DB, userEmail)
	if err != nil {
		fmt.Println("1")
		fmt.Println(err)
		helper.ErrorPage(w, 500)
		return
	}
	fmt.Println("======------======")
	fmt.Println(user)
	groupId := r.FormValue("groupId")
	title := html.EscapeString(strings.TrimSpace(r.FormValue("title")))
	des := html.EscapeString(strings.TrimSpace(r.FormValue("description")))
	date := r.FormValue("date")
	time := r.FormValue("time")
	var event = models.Event{}
	event.Userid = user.ID
	event.Title = title
	event.Description = des
	event.Date = date
	event.Time = time
	grId, errGrid := strconv.Atoi(groupId)
	if errGrid != nil {
		helper.ErrorPage(w, 500)
		return
	}

	event.GroupId = grId

	errc := event.CreateEvent(DB)
	if errc != nil {
		fmt.Println(errc)
		helper.ErrorPage(w, 500)
		return
	}

	helper.WriteJSON(w, 200, map[string]interface{}{"success": true}, nil)
}

func GetOption(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	_, _, UserID := helper.Auth(DB, r)

	eventid := r.URL.Query().Get("eventid")
	response := r.URL.Query().Get("response")

	eveId , erra := strconv.Atoi(eventid)
	if erra != nil {
		fmt.Println(0)
		fmt.Println(erra)
		return
	}
	
	option := models.Option{}
	option.EveID = eveId
	option.UserID = UserID
	
	if response == "Oui" {
		option.IsGoing = 1
		
	}else if response == "Non" {
		option.IsGoing = 0
	}else{
		
		return
	}

	erro := option.CreateOption(DB)
	if erro != nil {

		fmt.Println(erro)
		return
	}
	helper.WriteJSON(w,200,map[string]interface{}{"data":option},r.Header)

}