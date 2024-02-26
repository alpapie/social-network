package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"social_network/helper"
	"social_network/models"
	"sync"

	"github.com/gorilla/websocket"
)

type UserToNotif struct {
	conn   *websocket.Conn
	userID int
}

var (
	upgrader   websocket.Upgrader
	UsersNotif = make(map[int]*UserToNotif)
	mutex      = sync.RWMutex{}
)

func init() {
	upgrader = websocket.Upgrader{
		ReadBufferSize:  4096,
		WriteBufferSize: 4096,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}
}

func InitSocketNotification(w http.ResponseWriter, r *http.Request) {
	_, _, user_id := helper.Auth(DB, r)

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		helper.ErrorMessage(w, "impossible d'initialiser le websocket")
		return
	}

	user := &UserToNotif{conn: conn, userID: user_id}
	mutex.RLock()
	UsersNotif[user_id] = user
	mutex.RUnlock()

	fmt.Println(UsersNotif)
}

func GetNotification(w http.ResponseWriter, r *http.Request) {
	notif := models.Notification{}
	_, _, user_id := helper.Auth(DB, r)
	notifications, err := notif.GetNotification(DB, user_id)

	if err != nil {
		fmt.Println("notif error", err)
		helper.ErrorPage(w, 500)
		return
	}

	errjson := helper.WriteJSON(w, 200, map[string]interface{}{"success": true, "notifications": notifications}, nil)
	if errjson != nil {
		fmt.Println(errjson)
		helper.ErrorPage(w, 500)
		return
	}
}

func SendSocketNotification(notificationObject models.Notification, action string) {
	mutex.RLock()
	receiver, isOnline := UsersNotif[notificationObject.User_id]
	mutex.RUnlock()
	if isOnline {

		err:=receiver.conn.WriteMessage(websocket.TextMessage,toJSON(notificationObject,true))
		if err != nil {
			fmt.Println("Error sending socket notification", err)
			receiver.conn.WriteMessage(websocket.TextMessage,toJSON(models.Notification{},false))
			return
		}
	}
}

func toJSON(notificationObject models.Notification, success bool) []byte {
	jsonData, _ := json.Marshal(map[string]interface{}{
		"success": success,
		"data":    notificationObject,
	})
	return jsonData
}
