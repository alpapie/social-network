package controller

import (
	"fmt"
	"net/http"
	helper "social_network/helper"
	"social_network/models"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var Clients = make(map[*Client]bool)

func WsHandler(w http.ResponseWriter, r *http.Request) {

	_, _, UserID := helper.Auth(DB, r)
	user := models.User{ID: UserID}
	user.GetUserById(DB, UserID)

	fmt.Println("received a request YYYYYYYYYYYYYEEEEEEEEEESSSSSSSSSSSS")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	// defer conn.Close()
	client := &Client{Conn: conn, FirstName: user.FirstName, LastName: user.LastName, ID: user.ID}
	Clients[client] = true
	fmt.Println("NEW CLIENT ", client)
	// fmt.Println("the new client is ", client.Session.Username)
	// mes, _ := json.Marshal(ConnexionMsg{Action: "connexion", Username: user.FirstName+" "+user.LastName})
	// PreventAllUsers(mes)
	go client.Listen()

}
