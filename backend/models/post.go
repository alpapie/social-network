package models

import (
	"database/sql"
	"fmt"
	"time"

)

type Post struct {
	Id           int       `json:"id"`
	Group_id     int       `json:"Group_id"`
	User_id      int       `json:"User_id" `
	Titre        string    `json:"titre"`
	Image        string    `json:"image"`
	Content      string    `json:"content"`
	Privacy      string    `json:"privacy"`
	PostGroup    string    `json:"postGroup"`
	CreationDate time.Time `json:"creationDate"`
	AllowedUsers  []int
}

func (P *Post) Create(controllerDB *sql.DB) (int, error) {

	statement, err := controllerDB.Prepare("INSERT INTO Post (Group_id, User_id, titre, image, content, privacy, postGroup,creationDate) VALUES (?,?,?,?,?,?,?,?)")
	if err != nil {
		return -1, err
	}

	sqlResult, err := statement.Exec(P.Group_id, P.User_id, P.Titre, P.Image, P.Content, P.Privacy, P.PostGroup, time.Now())
	if err != nil {
		return -1, err
	}

	lastInsertedId, err := sqlResult.LastInsertId()
	if err != nil {
		return -1, err
	}
	var idPost = int(lastInsertedId)

	if (P.Privacy == "almostprivate") {
		er :=RegisterAllowedUsers(P.AllowedUsers , idPost , controllerDB)
		if er != nil {
			return -1, er
		}
	}

	return idPost , nil
}

func RegisterAllowedUsers(users []int , idPost int , controllerDB *sql.DB ) error {
	statement, err := controllerDB.Prepare("INSERT INTO AllowedPost (Post_id , User_id) VALUES (?,?)")
	if (err != nil ) {
		return err
	}

	for User_id := range users {
		_, er :=statement.Exec(idPost,User_id)
		if er != nil {
			return er
		}
	}
	return nil
}

func (P *Post) Check() {
	fmt.Println("the post " , P.User_id , P.User_id == 0)
}
