package models

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID         int
	FirstName  string
	LastName   string
	Email      string
	NickName   string
	BirthName  string
	Avatar     string
	AboutMe    string
	IsPublic   int
	TokenLogin string
}

func (u *User)GetUserById(db *sql.DB, id int) ( error) {
	query := `SELECT id, firstName, lastName, email, nickName, birthName, avatar, aboutMe, ispublic, tokenLogin FROM User WHERE id = ?`
	
	err := db.QueryRow(query, id).Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.NickName, &u.BirthName, &u.Avatar, &u.AboutMe, &u.IsPublic, &u.TokenLogin)
	if err != nil {
		return fmt.Errorf("GetUserById %d: %v", id, err)
	}

	return nil
}

func (u *User)CreateUser(db *sql.DB) error {
	query := `INSERT INTO User (firstName, lastName, email, nickName, birthName, avatar, aboutMe, ispublic, tokenLogin) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := db.Exec(query, u.FirstName, u.LastName, u.Email, u.NickName, u.BirthName, u.Avatar, u.AboutMe, u.IsPublic, u.TokenLogin)
	if err != nil {
		return fmt.Errorf("CreateU: %v", err)
	}

	return nil
}
