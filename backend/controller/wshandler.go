package controller

import (
	"fmt"
	"net/http"
	helper "social_network/helper"
	"social_network/models"
	"sync"

	"github.com/gorilla/websocket"
)

var upgraders = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var Clients = make(map[*Client]bool)
var clientsMutex = &sync.Mutex{}
var i = 0
var tabclient = []int{}

func WsHandler(w http.ResponseWriter, r *http.Request) {

	_, _, UserID := helper.Auth(DB, r)
	user := models.User{ID: UserID}
	user.GetUserById(DB, UserID)
	groups , err:= models.GetMemberGroupsByUserID(DB,UserID)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	fmt.Println("received a request YYYYYYYYYYYYYEEEEEEEEEESSSSSSSSSSSS")
	conn, err := upgraders.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	// defer conn.Close()

	client := &Client{Conn: conn, FirstName: user.FirstName, LastName: user.LastName, ID: user.ID}
	for _ , g := range groups {
		client.Groups = append(client.Groups, g.ID)
	}
	clientsMutex.Lock()
	Clients[client] = true
	clientsMutex.Unlock()
	fmt.Println("NEW CLIENT ", client)
	tabclient = append(tabclient, i)
	i++

	// fmt.Println("the new client is ", client.Session.Username)
	// mes, _ := json.Marshal(ConnexionMsg{Action: "connexion", Username: user.FirstName+" "+user.LastName})
	// PreventAllUsers(mes)
	fmt.Println("=========Online USERS======")
	fmt.Println(GetOnlineUsers())
	SendOnlineUsers()

	go client.Listen()

}
