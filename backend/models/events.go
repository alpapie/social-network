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
	sqlStatement := `INSERT INTO event (user_id, group_id, title, description, date, time) VALUES (?, ?, ?, ?, ?, ?)`
	_, err := db.Exec(sqlStatement, event.Userid, event.GroupId, event.Title, event.Description, event.Date, event.Time)
	return err
}

func GetEventsByGroupId(db *sql.DB, groupId int) ([]Event, error) {
	var events []Event
	sqlStatement := `SELECT id, user_id, group_id, title, description, date, time FROM event WHERE group_id = ?`
	rows, err := db.Query(sqlStatement, groupId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var e Event
		err := rows.Scan(&e.ID, &e.Userid, &e.GroupId, &e.Title, &e.Description, &e.Date, &e.Time)
		if err != nil {
			return nil, err
		}
		events = append(events, e)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}
