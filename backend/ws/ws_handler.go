package ws

import (
	"encoding/json"
	"log"
	"net/http"
	"sort"
	"time"

	"github.com/gorilla/websocket"
)

type Handler struct {
	hub *Hub
}

func NewHandler(h *Hub) *Handler {
	return &Handler{hub: h}
}

type CreateRoomReq struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func writeResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) CreateRoom(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req CreateRoomReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, ok := h.hub.Rooms[req.ID]; ok {
		http.Error(w, "A room with the same ID already exists", http.StatusConflict)
		return
	}

	h.hub.Rooms[req.ID] = &Room{
		ID:      req.ID,
		Name:    req.Name,
		Clients: map[string]*Client{},
	}

	json.NewEncoder(w).Encode(req)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (h *Handler) JoinRoom(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to upgrade connection: %v", err)
		http.Error(w, "Failed to upgrade connection", http.StatusBadRequest)
		return
	}

	roomId := r.URL.Query().Get("roomId")
	ClientId := r.URL.Query().Get("userId")
	username := r.URL.Query().Get("username")

	cl := &Client{
		Conn:     conn,
		Message:  make(chan *Message),
		ID:       ClientId,
		RoomID:   roomId,
		Username: username,
	}

	m := &Message{
		Content:   "NEW COMER",
		RoomID:    roomId,
		Username:  username,
		Action:    "connexion",
		SenderID:  ClientId,
		TimeStamp: time.Now().Format("Jan 2, 2006 at 3:04pm"),
	}

	h.hub.Register <- cl
	h.hub.Broadcast <- m
	go cl.writeMessage()
	cl.readMessage(h.hub)
}

type RoomRes struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) GetRooms(w http.ResponseWriter, r *http.Request) {
	rooms := make([]RoomRes, 0)

	for _, r := range h.hub.Rooms {
		rooms = append(rooms, RoomRes{
			ID:   r.ID,
			Name: r.Name,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	writeResponse(w, http.StatusOK, rooms)
}

type ClientRes struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

func (h *Handler) GetClients(w http.ResponseWriter, r *http.Request) {
	var clients []ClientRes

	roomId := r.URL.Query().Get("roomId")

	if _, ok := h.hub.Rooms[roomId]; !ok {
		clients = make([]ClientRes, 0)
		writeResponse(w, http.StatusOK, clients)
		return
	}
	for _, c := range h.hub.Rooms[roomId].Clients {
		clients = append(clients, ClientRes{
			ID:       c.ID,
			Username: c.Username,
		})
	}

	sort.Slice(clients, func(i, j int) bool {
		return clients[i].Username < clients[j].Username
	})

	writeResponse(w, http.StatusOK, clients)
}
