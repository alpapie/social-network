package models

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Notification struct {
	ID       int    `json:"id"`
	User_id  int    `json:"user_id"`
	SenderID int    `json:"sender_id"`
	Type     string `json:"type"`
	Status   string `json:"status"`
}

func (n *Notification) CreateNotification(db *sql.DB) error {

	checkQuery := `SELECT COUNT(*) FROM Notfication WHERE User_id = ? AND send_id = ? AND Type = ?`
	row := db.QueryRow(checkQuery, n.User_id, n.SenderID, n.Type)
	var count int
	err := row.Scan(&count)
	if err != nil {
		return fmt.Errorf("failed to check for existing notification: %v", err)
	}
	if count > 0 {
		return fmt.Errorf("this notif already done")
	}

	query := `INSERT INTO Notfication (User_id, send_id, type, status) VALUES (?, ?, ?, ?)`
	_, err1 := db.Exec(query, n.User_id, n.SenderID, n.Type, n.Status)
	if err1 != nil {
		return fmt.Errorf("failed to create notification: %v", err1)
	}
	return nil
}
