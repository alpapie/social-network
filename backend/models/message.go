package models

import (
	"database/sql"
	"fmt"
)

type Message struct {
	Id           int    `json:"id"`
	Sender_id    int `json:"sender_id"`
	Receiver_id  int `json:"receiver_id"`
	Content      string `json:"content"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	CreationDate string `json:"creationDate"`
	GroupId int `json:"groupId"`
}

func (m *Message) Create(controllerDB *sql.DB) (int , error){
	statement, err := controllerDB.Prepare("INSERT INTO message (sender_id, receiver_id, content) VALUES (?,?,?)")
	if err != nil {
		return -1, err
	}

	sqlResult, err := statement.Exec(m.Sender_id, m.Receiver_id, m.Content)
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

func GetDiscussion(db *sql.DB, receiverID int, senderId int) ([]Message, error) {
	stmt, err := db.Prepare(`
	SELECT  M.id , U.firstname, U.lastname , M.sender_id ,M.receiver_id , M.content , M.dateCreation from message as M 
	join User as U on U.id = M.sender_id
	where (M.sender_id = ? and M.receiver_id = ?) or (M.receiver_id =? and M.sender_id  = ?) 
			order by M.dateCreation asc;
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare discussion statement: %v", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(senderId, receiverID ,senderId , receiverID)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var messages []Message
	for rows.Next() {
		var m Message
		err := rows.Scan(&m.Id, &m.FirstName ,&m.LastName, &m.Sender_id, &m.Receiver_id, &m.Content, &m.CreationDate)
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

