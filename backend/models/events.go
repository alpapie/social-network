package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Event struct {
	ID          int    `json:"id"`
	Userid      int    `json:"user_id"`
	GroupId     int    `json:"group_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Time        string `json:"time"`
}

func (event *Event) CreateEvent(db *sql.DB) error {
	sqlStatement := `INSERT INTO events (user_id, group_id, title, description, date, time) VALUES (?, ?, ?, ?, ?, ?)`
	_, err := db.Exec(sqlStatement, event.Userid, event.GroupId, event.Title, event.Description, event.Date, event.Time)
	return err
}
