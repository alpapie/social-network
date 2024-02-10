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

func GetNotificationByUserIDAndType(db *sql.DB, SenderID int, userID int, notificationType string) ([]Notification, error) {
	stmt, err := db.Prepare(`
        SELECT User_id, send_id, Type, Status
        FROM Notfication
        WHERE send_id = ? AND Type = ? AND User_id = ?
    `)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare get notification by user ID and type statement: %v", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(SenderID, notificationType, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var notifications []Notification
	for rows.Next() {
		var n Notification
		err := rows.Scan(&n.User_id, &n.SenderID, &n.Type, &n.Status)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		notifications = append(notifications, n)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration: %v", err)
	}

	return notifications, nil
}
