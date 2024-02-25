package controller

import (
	"encoding/json"
	"fmt"
	"social_network/models"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn      *websocket.Conn
	ID        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	group_id int `json:"groupid"`
}
type ConnexionMsg struct {
	Action   string
	Username string
}

func (c *Client) Listen() {
	var newMessage models.Message
	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			DeleteUser(c.ID)
			break
		}
		// mes , _ := json.Marshal(ConnexionMsg{Action: "createpost" , Username: c.Session.Username})
		json.Unmarshal(message, &newMessage)
		newMessage.FirstName = c.FirstName
		newMessage.LastName = c.LastName
		newMessage.Sender_id = c.ID
		newMessage.CreationDate = time.Now().Format("Jan 2, 2006 at 3:04pm")
		fmt.Println("new message: ", newMessage)
		// fmt.Println("the error :", erro, "new message  ", newMessage.Action, newMessage.Content)
		_ , Erm := newMessage.Create(DB)
		if Erm != nil {
			return
		}
		PreventSelectedUser(newMessage)
		fmt.Println("get private message ", newMessage, "bytes", message)

	}
}

func (c *Client) ListenGroup() {
	var newMessage models.Message
	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			DeleteUser(c.ID)
			break
		}
		// mes , _ := json.Marshal(ConnexionMsg{Action: "createpost" , Username: c.Session.Username})
		json.Unmarshal(message, &newMessage)
		newMessage.FirstName = c.FirstName
		newMessage.LastName = c.LastName
		newMessage.Sender_id = c.ID
		newMessage.CreationDate = time.Now().Format("Jan 2, 2006 at 3:04pm")
		newMessage.GroupId = c.group_id
		fmt.Println("new message: ", newMessage)
		// fmt.Println("the error :", erro, "new message  ", newMessage.Action, newMessage.Content)
		// _ , Erm := newMessage.Create(DB)
		// if Erm != nil {
		// 	return
		// }
		PreventGroupMembers(newMessage)
		fmt.Println("get private message ", newMessage, "bytes", message)

	}
}

func PreventAllUsers(message []byte) {
	for c := range Clients {
		c.Conn.WriteMessage(websocket.TextMessage, message)
	}
}

func PreventGroupMembers(message models.Message) {
    for c, connected := range Clients {
        if connected && c.group_id == message.GroupId {
            byteMes, _ := json.Marshal(message)
            c.Conn.WriteMessage(websocket.TextMessage, byteMes)
        }
    }
}

func PreventSelectedUser(message models.Message) {
	for c, connected := range Clients {
		if connected && (c.ID == message.Receiver_id || c.ID == message.Sender_id)  {
			byteMes, _ := json.Marshal(message)

			c.Conn.WriteMessage(websocket.TextMessage, byteMes)
		}
	}
}

func DeleteUser(id int) {
	for c, connected := range Clients {
		if c.ID == id && connected {
			delete(Clients, c)
			break
		}
	}
}
