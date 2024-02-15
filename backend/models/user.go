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
	Pasword    string
	NickName   string
	BirthDate  string
	Avatar     string
	AboutMe    string
	IsPublic   int
	TokenLogin string
}

func (u *User) GetUserById(db *sql.DB, id int) error {
	query := `SELECT id, firstName, lastName, email, nickName, birthDate, avatar, aboutMe, ispublic FROM User WHERE id = ?`

	err := db.QueryRow(query, id).Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.NickName, &u.BirthDate, &u.Avatar, &u.AboutMe, &u.IsPublic)
	if err != nil {
		return fmt.Errorf("GetUserById %d: %v", id, err)
	}

	return nil
}

func (u *User) CreateUser(db *sql.DB) error {
	// Vérifier si l'email existe déjà
	var count int
	query := `SELECT COUNT(*) FROM User WHERE email = ?`
	err := db.QueryRow(query, u.Email).Scan(&count)
	if err != nil {
		return fmt.Errorf("failed to check email existence: %v", err)
	}
	if count > 0 {
		return fmt.Errorf("email already exists")
	}

	query = `INSERT INTO User (firstName, lastName, email, nickName, birthDate, avatar, aboutMe, ispublic, tokenLogin) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err = db.Exec(query, u.FirstName, u.LastName, u.Email, u.NickName, u.BirthDate, u.Avatar, u.AboutMe, u.IsPublic, u.TokenLogin)
	if err != nil {
		return fmt.Errorf("CreateU: %v", err)
	}

	return nil
}

func (u *User) GetUserByToken(db *sql.DB, token string) error {
	query := `SELECT id, firstName, lastName, email, nickName, birthDate, avatar, aboutMe, ispublic FROM User WHERE tokenLogin = ?`

	err := db.QueryRow(query, token).Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.NickName, &u.BirthDate, &u.Avatar, &u.AboutMe, &u.IsPublic)
	if err != nil {
		return fmt.Errorf("GetUserById %s: %v", token, err)
	}

	return nil
}

func (u *User) AddFlow(db *sql.DB,follow_id int)(error){
	req:=`insert INTO "Follow" ("User_id","Follower_id") VALUES(?,?)`
	_,err:=db.Exec(req,u.ID,follow_id,)
	return err
}