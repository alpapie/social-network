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

func (u *User) IsGroupmemeber(db *sql.DB, groupID int) (bool, error) {

	var count int
	query := `SELECT COUNT(*) FROM joinner WHERE User_id = ? AND Group_id = ?`
	err := db.QueryRow(query, u.ID, groupID).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("failed to check group member: %v", err)
	}
	if count < 1 {
		return false, fmt.Errorf("not a member")
	}
	fmt.Println(count)

	return true, nil
}

func (u *User) GetUserById(db *sql.DB, id int) error {
	query := `SELECT id, firstName, lastName, email, nickName, birthDate, avatar, aboutMe, ispublic FROM User WHERE id = ?`

	err := db.QueryRow(query, id).Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.NickName, &u.BirthDate, &u.Avatar, &u.AboutMe, &u.IsPublic)
	if err != nil {
		return fmt.Errorf("GetUserById %d: %v", id, err)
	}

	return nil
}

func (u *User) GetUserByEmail(db *sql.DB, email string) error {
	query := `SELECT id, firstName, lastName, email, nickName, birthDate, avatar, aboutMe, ispublic FROM User WHERE email = ?`

	err := db.QueryRow(query, email).Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.NickName, &u.BirthDate, &u.Avatar, &u.AboutMe, &u.IsPublic)
	if err != nil {
		return fmt.Errorf("GetUserById %s: %v", email, err)
	}

	return nil
}

func (u *User) CreateUser(db *sql.DB) error {
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

func (u *User) AddFollower(db *sql.DB, follow_id int) error {
	req := `insert INTO "Follow" ("User_id","Follower_id") VALUES(?,?)`
	_, err := db.Exec(req, u.ID, follow_id)
	return err
}

func (u *User) GetFollowed(db *sql.DB) ([]User, error) {
	query := `SELECT User_id FROM Follow WHERE Follower_id = ?`
	rows, err := db.Query(query, u.ID)
	if err != nil {
		return nil, fmt.Errorf("GetFollowed: %v", err)
	}
	defer rows.Close()

	var followers []User
	for rows.Next() {
		var followerID int
		if err := rows.Scan(&followerID); err != nil {
			return nil, fmt.Errorf("GetFollowed scan: %v", err)
		}
		follower := User{}
		if err := follower.GetUserById(db, followerID); err != nil {
			return nil, fmt.Errorf("GetFollowed GetUserById: %v", err)
		}
		followers = append(followers, follower)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetFollowed rows.Err: %v", err)
	}

	return followers, nil
}

func (u *User) GetFollowers(db *sql.DB) ([]User, error) {
	query := `SELECT Follower_id FROM Follow WHERE User_id = ?`
	rows, err := db.Query(query, u.ID)
	if err != nil {
		return nil, fmt.Errorf("GetFollower: %v", err)
	}
	defer rows.Close()

	var followed []User
	for rows.Next() {
		var followedID int
		if err := rows.Scan(&followedID); err != nil {
			return nil, fmt.Errorf("GetFollower scan: %v", err)
		}
		followedUser := User{}
		if err := followedUser.GetUserById(db, followedID); err != nil {
			return nil, fmt.Errorf("GetFollower GetUserById: %v", err)
		}
		followed = append(followed, followedUser)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetFollower rows.Err: %v", err)
	}

	return followed, nil
}
