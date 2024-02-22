package ws

import (
	"database/sql"
	"log"
	"social_network/models"
	"strconv"
)

type Room struct {
	ID      string             `json:"id"`
	Name    string             `json:"name"`
	Clients map[string]*Client `json:"clients"`
}

type Hub struct {
	Rooms      map[string]*Room
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *Message
}

func NewHub(DB *sql.DB) *Hub {
	hub := &Hub{
		Rooms:      make(map[string]*Room),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *Message, 100),
	}

	// Appelez GetAllGroups et ajoutez chaque groupe à la structure Rooms
	groups, err := models.GetAllGroups(DB)
	if err != nil {
		log.Printf("Error getting all groups: %v", err)
		return hub
	}

	for _, group := range groups {
		roomID := strconv.Itoa(group.ID)
		hub.Rooms[roomID] = &Room{
			ID:      roomID,
			Name:    group.Title,
			Clients: make(map[string]*Client),
		}
	}

	return hub
}

func (h *Hub) Run() {
	for {
		select {
		case cl := <-h.Register:
			if _, ok := h.Rooms[cl.RoomID]; ok {
				r := h.Rooms[cl.RoomID]

				if _, ok := r.Clients[cl.ID]; !ok {
					r.Clients[cl.ID] = cl
				}
			}
		case cl := <-h.Unregister:
			if _, ok := h.Rooms[cl.RoomID]; ok {
				if _, ok := h.Rooms[cl.RoomID].Clients[cl.ID]; ok {
					if len(h.Rooms[cl.RoomID].Clients) != 0 {
						h.Broadcast <- &Message{
							Content:  "left the chat",
							RoomID:   cl.RoomID,
							Username: cl.Username,
							SenderID: cl.ID,
							Action:   "deconnexion",
						}
					}

					delete(h.Rooms[cl.RoomID].Clients, cl.ID)
					close(cl.Message)
				}
			}

		case m := <-h.Broadcast:
			if _, ok := h.Rooms[m.RoomID]; ok {
				if m.Action == "message" || m.Action == "typing" {
					if m.Action != "typing" {
						sender, senderOk := h.Rooms[m.RoomID].Clients[m.SenderID]
						if senderOk {
							select {
							case sender.Message <- m:
							default:
								log.Printf("Failed to send message to client %s", sender.ID)
							}
						}
					}
					recipient, recipientOk := h.Rooms[m.RoomID].Clients[m.RecipientID]
					if recipientOk {
						select {
						case recipient.Message <- m:
						default:
							log.Printf("Failed to send message to client %s", recipient.ID)
						}
					}
				} else {
					for _, client := range h.Rooms[m.RoomID].Clients {
						select {
						case client.Message <- m:
						default:
							log.Printf("Failed to send message to client %s", client.ID)
						}
					}
				}
			}
		}
	}
}
