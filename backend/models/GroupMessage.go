package models

import (
	"database/sql"
	"fmt"
)

type GroupMessage struct {
	Id           int    `json:"id"`
	User_id    int `json:"sender_id"`
	Group_id  int `json:"receiver_id"`
	Content      string `json:"content"`
	DateCreation string `json:"datecreation"`
}

func (gm *GroupMessage) CreateGroupmesage(controllerDB *sql.DB) (int , error){
	statement, err := controllerDB.Prepare("INSERT INTO groupmessage (user_id, group_id, content) VALUES (?,?,?)")
	if err != nil {
		return -1, err
	}

	sqlResult, err := statement.Exec(gm.User_id, gm.Group_id, gm.Content)
	if err != nil {
		return -1, err
	}

	lastInsertedId, err := sqlResult.LastInsertId()
	if err != nil {
		return -1, err
	}
	var idmessage = int(lastInsertedId)

	return idmessage, nil
}

func GetGroupDiscussion(db *sql.DB, groupID int) ([]Message, error) {
	stmt, err := db.Prepare(`
	SELECT  G.id , U.firstname, U.lastname , G.User_id  , G.content , G.dateCreation
		from GroupMessage as G JOIN user as U on G.User_id = U.id 
		WHERE G.group_id = ?
		order by G.dateCreation asc;
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare discussion statement: %v", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(groupID)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var messages []Message
	for rows.Next() {
		var m Message
		err := rows.Scan(&m.Id, &m.FirstName ,&m.LastName, &m.Sender_id, &m.Content, &m.CreationDate)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		messages = append(messages, m)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration: %v", err)
	}

	return messages, nil
}