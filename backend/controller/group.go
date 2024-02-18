package controller

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
	helper "social_network/helper"
	"social_network/models"
	"strconv"
	"strings"
)

type NewGroup struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	NotifStatus bool      `json:"notif_type"`
	SuggestList []NewUser `json:"suggests"`
}
type NewUser struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Avatar      string `json:"avatar"`
	Isrequested bool   `json:"is_requested"`
}
type GroupDetail struct {
	NbrFollowers int            `json:"nbrfollowers"`
	Events       []models.Event `json:"events"`
	Groupdata    models.Group   `json:"groupdata"`
}

func CreateFollowGroup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HERE")
	type GroupFollowData struct {
		GroupID int `json:"groupId"`
		UserID  int `json:"userId"`
	}

	var followData GroupFollowData

	err := json.NewDecoder(r.Body).Decode(&followData)
	if err != nil {
		fmt.Println("ggg")
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("userID", followData.UserID)
	fmt.Println("groupID", followData.GroupID)
	var user = models.User{}

	errr := user.GetUserById(DB, followData.UserID)
	if errr != nil {
		fmt.Println("111")
		return
	}
	fmt.Println("THE USER")
	fmt.Println(user)
	sess, err1 := user.HasActiveSession(DB)
	if err1 != nil || !sess {
		fmt.Println(err1)
		fmt.Println("ff")
		return
	}
	fmt.Println(sess)
	group, errg := models.GetGroupByID(DB, followData.GroupID)
	if errg != nil {
		fmt.Println("t")
		return
	}

	var notification = models.Notification{}
	notification.SenderID = user.ID
	notification.User_id = group.UserID
	notification.Type = "follow-Group"
	notification.Group_id = group.ID
	notification.Status = "false"
	ern := notification.CreateNotification(DB)
	if ern != nil {
		fmt.Println(ern)
	}

}

func CreateGroupHandler(w http.ResponseWriter, r *http.Request) {

	auth, userEmail ,_:= helper.Auth(DB, r)
	if !auth {
		fmt.Println("Not registered")
		return
	}
	var user = models.User{}
	err := user.GetUserByEmail(DB, userEmail)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var newGroup models.Group
	fmt.Println(r)
	err = json.NewDecoder(r.Body).Decode(&newGroup)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		fmt.Println("in decode")
		return
	}
	if newGroup.Title == "" || newGroup.Description == "" {
		http.Error(w, "Title and description cannot be empty", http.StatusBadRequest)
		return
	}
	newGroup.UserID = user.ID
	lastInsertId, err := newGroup.CreateGroup(DB)
	if err != nil {
		http.Error(w, "Failed to create group", http.StatusInternalServerError)
		fmt.Println(err)
		fmt.Println("in user")
		return
	}
	er := models.JoinGroup(DB, user.ID, int(lastInsertId))
	if er != nil {
		fmt.Println(er)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "success",
		"message": "Group created successfully",
		"groupId": lastInsertId,
	})
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	var user = models.User{}
	auth, userEmail ,_:= helper.Auth(DB, r)
	if !auth {
		fmt.Println("01")
		return
	}
	err := user.GetUserByEmail(DB, userEmail)
	if err != nil {
		fmt.Println("1")
		fmt.Println(err)
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
		helper.ErrorPage(w, 400)
		return
	}

	event.GroupId = grId

	errc := event.CreateEvent(DB)
	if errc != nil {
		fmt.Println(errc)
		helper.ErrorPage(w, 400)
		return
	}

	helper.WriteJSON(w, 200, map[string]interface{}{"success": true}, nil)
}

func GetGroupDetail(w http.ResponseWriter, r *http.Request) {
	var user = models.User{}
	_, userEmail,_ := helper.Auth(DB, r)
	err := user.GetUserByEmail(DB, userEmail)
	if err != nil {
		return
	}

	groupId := strings.Split(r.URL.Path, "/")

	grId, errg := strconv.Atoi(groupId[len(groupId)-1])
	if errg != nil {
		return
	}
	gr, err := models.GetGroupByID(DB, grId)
	if err != nil {
		helper.ErrorPage(w, 500)
		return
	}
	ismem, errm := user.IsGroupmemeber(DB, grId)
	if errm != nil || !ismem {
		helper.ErrorPage(w, 500)
		return
	}
	events, erre := models.GetEventsByGroupId(DB, grId)
	if erre != nil {
		helper.ErrorPage(w, 500)
		return
	}
	nbr, errgm := models.GetGroupMemberCount(DB, grId)
	if errgm != nil {
		helper.ErrorPage(w, 500)
		return
	}
	groupapagedata := GroupDetail{}
	groupapagedata.NbrFollowers = nbr
	groupapagedata.Events = events
	groupapagedata.Groupdata = *gr

	helper.WriteJSON(w, 200, map[string]interface{}{"success": true, "groupDetail": groupapagedata}, nil)
}
