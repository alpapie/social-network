package global

import (
	"database/sql" // Assurez-vous que ce chemin correspond au nom de votre module
	"social_network/ws"
)

var (
	HUB        *ws.Hub     // Déclaration de la variable globale hub
	WS_HANDLER *ws.Handler // Déclaration de la variable globale wsHandler
)

func Init(db *sql.DB) {
	HUB = ws.NewHub(db)
	WS_HANDLER = ws.NewHandler(HUB)
	go HUB.Run()
}