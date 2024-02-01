package models

import (
	"database/sql"
	"time"
)

type Post struct {
	Id           int       `json:"id"`
	Group_id     int       `json:"Group_id"`
	User_id      int       `json:"User_id" `
	Titre        string    `json:"titre"`
	Image        string    `json:"image"`
	Content      string    `json:"content"`
	Privaty      string    `json:"privaty"`
	PostGroup    string    `json:"postGroup"`
	CreationDate time.Time `json:"creationDate"`
}

func (P *Post) Create(controllerDB *sql.DB) (int, error) {

	statement, err := controllerDB.Prepare("INSERT INTO Post (Group_id, User_id, titre, image, content, privaty, postGroup,creationDate) VALUES (?,?,?,?,?,?,?,?)")
	if err != nil {
		return -1, err
	}
	defer statement.Close()

	sqlResult, err := statement.Exec(P.Group_id, P.User_id, P.Titre, P.Image, P.Content, P.Privaty, P.PostGroup, time.Now())
	if err != nil {
		return -1, err
	}

	lastInsertedId, err := sqlResult.LastInsertId()
	if err != nil {
		return -1, err
	}

	return int(lastInsertedId), nil
}
