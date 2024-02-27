package models

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Notification struct {
	ID         int    `json:"id"`
	User_id    int    `json:"user_id"`
	SenderID   int    `json:"sender_id"`
	FirstName  string `json:"firstname"`
	LastName   string `json:"lastname"`
	Avatar     string `json:"avatar"`
	GroupTitle string `json:"grouptitle"`
	Group_id   int    `json:"group_id"`
	Type       string `json:"type"`
	Status     string `json:"status"`
}

func (n *Notification) CreateNotification(db *sql.DB) error {
	query := `INSERT INTO Notification (User_id, send_id, type, status, group_id) VALUES (?, ?, ?, ?, ?)`
	id, err1 := db.Exec(query, n.User_id, n.SenderID, n.Type, "0", n.Group_id)
	if err1 != nil {
		return fmt.Errorf("failed to create notification: %v", err1)
	}
	notifId,err:=id.LastInsertId()
	n.ID=int(notifId)
	return err
}

func (N *Notification) GetNotificationByUserIDAndTypeAndnotifid(db *sql.DB, notifid, SenderID, userID, groupID int) error {
	stmt, err := db.Prepare(`
	SELECT notif.id,notif.status, notif.type, notif.send_id , coalesce(g.id, 0), u."firstName", u."lastName",u.avatar, coalesce(g.title,"") from "User" as u INNER join "Notification" as notif on notif.send_id=u.id LEFT join "Group" as g on g.id=notif."Group_id" WHERE 
	notif."User_id"=? and notif.id=? and notif.send_id=? and notif."Group_id"=?
    `)
	if err != nil {
		return fmt.Errorf("failed to prepare get notification by user ID and type statement: %v", err)
	}
	defer stmt.Close()
	return stmt.QueryRow(userID, notifid, SenderID, groupID).Scan(&N.ID, &N.Status, &N.Type, &N.SenderID, &N.Group_id, &N.FirstName, &N.LastName, &N.Avatar, &N.GroupTitle)
}

func (n Notification) MarkAsRead(db *sql.DB) error {
	req := `UPDATE Notification SET status = 1 WHERE id=?`
	_, err := db.Exec(req, n.ID)
	if err != nil {
		return fmt.Errorf("error when marking as read: %v", err)
	}
	return nil
}

func GetNotificationByUserIDAndType(db *sql.DB, SenderID int, userID int, notificationType string, groupID int) ([]Notification, error) {
	stmt, err := db.Prepare(`
        SELECT User_id, send_id, Type, Status
        FROM Notification
        WHERE send_id = ? AND Type = ? AND User_id = ? AND Group_id = ?
    `)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare get notification by user ID and type statement: %v", err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(SenderID, notificationType, userID, groupID)
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

func (N Notification) GetNotification(db *sql.DB, user_id int) ([]Notification, error) {
	req := `SELECT notif.id,notif.status, notif.type, notif.send_id , coalesce(g.id, 0), u."firstName", u."lastName",u.avatar, coalesce(g.title,"") from "User" as u INNER join "Notification" as notif on notif.send_id=u.id LEFT join "Group" as g on g.id=notif."Group_id" WHERE notif."User_id"=? and notif.status!=1 `
	row, err := db.Query(req, user_id)
	notifications := []Notification{}
	if err != nil {
		return notifications, err
	}

	for row.Next() {
		row.Scan(&N.ID, &N.Status, &N.Type, &N.SenderID, &N.Group_id, &N.FirstName, &N.LastName, &N.Avatar, &N.GroupTitle)
		notifications = append(notifications, N)
	}
	return notifications, nil
}

func (notif *Notification) GetNotifById(DB *sql.DB, id int) error {
	req := `selec * from "Notification" where id=?`
	err := DB.QueryRow(req, id).Scan(&notif.ID, &notif.User_id, &notif.SenderID, &notif.Group_id, &notif.Type, &notif.Status)
	return err
}
