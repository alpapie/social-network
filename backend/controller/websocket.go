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
	Groups []int `json:"groupid"`
}
type ConnexionMsg struct {
	Action   string
	Username string
}
type OnlineUsersMessage struct {
    Action   string    `json:"action"`
    Users    []models.User `json:"users"`
}

func (c *Client) Listen() {
	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			DeleteUser(c.ID)
			SendOnlineUsers()
			fmt.Println("UN USER A QUITTER LE WEBSOCKET")
			break
		}
		// mes , _ := json.Marshal(ConnexionMsg{Action: "createpost" , Username: c.Session.Username})
		var newMessage models.Message
		json.Unmarshal(message, &newMessage)
		newMessage.FirstName = c.FirstName
		newMessage.LastName = c.LastName
		newMessage.Sender_id = c.ID
		newMessage.CreationDate = time.Now().Format("Jan 2, 2006 at 3:04pm")
		fmt.Println("new message: ", newMessage)
		// fmt.Println("the error :", erro, "new message  ", newMessage.Action, newMessage.Content)
		if newMessage.GroupId == 0 {
			_ , Erm := newMessage.Create(DB)
			if Erm != nil {
				return
			}
			PreventSelectedUser(newMessage)
			fmt.Println("get private message ", newMessage, "bytes", message)
		} else {
			_ , Erm := newMessage.CreateGroupmesage(DB)
			if Erm != nil {
				return
			}
			PreventGroupMembers(newMessage)
		}
		

	}
}
func GetOnlineUsers() []models.User {
    var onlineUsers []models.User
    clientsMutex.Lock()
    for client, connected := range Clients {
        if connected {
			user := models.User{}
			err := user.GetUserById(DB, client.ID)
			if err != nil {
				return nil
			}
			
			onlineUsers = append(onlineUsers, user)
			
            
        }
    }
    clientsMutex.Unlock()
    return onlineUsers
}
func SendOnlineUsers() {
    onlineUsers := GetOnlineUsers()
    message := OnlineUsersMessage{
        Action: "onlineUsers",
        Users:  onlineUsers,
    }

    for client, connected := range Clients {
        if connected {
            err := client.Conn.WriteJSON(message)
            if err != nil {
                fmt.Printf("Error sending online users to client %d: %v\n", client.ID, err)
                DeleteUser(client.ID)
				SendOnlineUsers()
            }
        }
    }
}

// func (c *Client) ListenGroup() {
// 	var newMessage models.Message
// 	for {
// 		_, message, err := c.Conn.ReadMessage()
// 		if err != nil {
// 			DeleteUser(c.ID)
// 			break
// 		}
// 		// mes , _ := json.Marshal(ConnexionMsg{Action: "createpost" , Username: c.Session.Username})
// 		json.Unmarshal(message, &newMessage)
// 		newMessage.FirstName = c.FirstName
// 		newMessage.LastName = c.LastName
// 		newMessage.Sender_id = c.ID
// 		newMessage.CreationDate = time.Now().Format("Jan 2, 2006 at 3:04pm")
// 		newMessage.GroupId = c.group_id
// 		fmt.Println("new message: ", newMessage)
// 		// fmt.Println("the error :", erro, "new message  ", newMessage.Action, newMessage.Content)
// 		// _ , Erm := newMessage.Create(DB)
// 		// if Erm != nil {
// 		// 	return
// 		// }
// 		PreventGroupMembers(newMessage)
// 		fmt.Println("get private message ", newMessage, "bytes", message)

// 	}
// }

// func PreventAllUsers(message []byte) {
// 	for c := range Clients {
// 		c.Conn.WriteMessage(websocket.TextMessage, message)
// 	}
// }

func PreventGroupMembers(message models.Message) {
    for c, connected := range Clients {
        if connected && Contain(c.Groups ,message.GroupId){
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
	clientsMutex.Lock()
	for c, connected := range Clients {
		if c.ID == id && connected {
			delete(Clients, c)
			break
		}
	}
	clientsMutex.Unlock()
}

func Contain(tab []int , id int) bool {
	for _, t := range tab {
		if t == id {
			return true
		}
	}
	return false
}