package models

import (
	"database/sql"
	"fmt"
)

type Option struct {
	ID      int `json:"id"`
	EveID   int `json:"Eve_id"`
	UserID  int `json:"User_id"`
	IsGoing int `json:"isGoing"`
}

func (o *Option) CreateOption(db *sql.DB) error {

	query := `INSERT INTO option (eve_id, user_id, isGoing) VALUES (?, ?, ?)`
	_, err := db.Exec(query, o.EveID, o.UserID, o.IsGoing)
	if err != nil {
		return fmt.Errorf("CreateO: %v", err)
	}

	return nil
}

func GetAllOptions(db *sql.DB) ([]Option, error) {
    query := `SELECT * FROM option`
    rows, err := db.Query(query)
    if err != nil {
        return nil, fmt.Errorf("GetAllOptions: %v", err)
    }
    defer rows.Close()

    var options []Option
    for rows.Next() {
        var option Option
        err := rows.Scan(&option.ID, &option.EveID, &option.UserID, &option.IsGoing)
        if err != nil {
            return nil, fmt.Errorf("GetAllOptions: %v", err)
        }
        options = append(options, option)
    }

    if err = rows.Err(); err != nil {
        return nil, fmt.Errorf("GetAllOptions: %v", err)
    }

    return options, nil
}