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
	IsGoing *bool `json:"isgoing"`
}

func (event *Event) CreateEvent(db *sql.DB) error {
	sqlStatement := `INSERT INTO event (user_id, group_id, title, description, date, time) VALUES (?, ?, ?, ?, ?, ?)`
	_, err := db.Exec(sqlStatement, event.Userid, event.GroupId, event.Title, event.Description, event.Date, event.Time)
	return err
}

func GetEventsByGroupId(db *sql.DB, groupId int) ([]Event, error) {
    var events []Event
    sqlStatement := `SELECT e.id, e.user_id, e.group_id, e.title, e.description, e.date, e.time, o.isgoing FROM event as e LEFT JOIN Option as o on o.eve_id = e.id WHERE e.group_id = ?`
    rows, err := db.Query(sqlStatement, groupId)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var e Event
        var isGoing *bool
        err := rows.Scan(&e.ID, &e.Userid, &e.GroupId, &e.Title, &e.Description, &e.Date, &e.Time, &isGoing)
        if err != nil {
            return nil, err
        }
        e.IsGoing = isGoing
        events = append(events, e)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return events, nil
}

