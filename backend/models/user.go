package models

import (
	"database/sql"
	"fmt"
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
	Online     bool
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

func AreFriend(db *sql.DB, userId1 int, userId2 int) (bool, error) {
	var count int
	query := `SELECT count(F.id)from Follow as F  where (F.User_id = ? and F.Follower_id = ?) or
		(F.Follower_id = ?  and F.User_id = ?)`
	err := db.QueryRow(query, userId1, userId2, userId1, userId2).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("failed to check friendship: %v", err)
	}
	if count < 1 {
		return false, nil
	}
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
	if Isfollower(db, u.ID, follow_id) {
		return fmt.Errorf("already followed")
	}

	req := `insert INTO "Follow" ("User_id","Follower_id") VALUES(?,?)`
	_, err := db.Exec(req, u.ID, follow_id)
	return err
}

func (U *User) UnFollowUser(db *sql.DB, userToUnfollowId int) error {
	if !Isfollower(db, userToUnfollowId, U.ID) {
		return fmt.Errorf("not following")
	}
	req := `DELETE FROM "Follow" where User_id=? AND Follower_id=?`
	_, err := db.Exec(req, userToUnfollowId, U.ID)
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

func (u *User) GetFollowerANDFollowed(db *sql.DB) ([]User, error) {
	query := `
        SELECT User_id AS id FROM Follow WHERE Follower_id = ?
        UNION
        SELECT Follower_id AS id FROM Follow WHERE User_id = ?
    `
	rows, err := db.Query(query, u.ID, u.ID)
	if err != nil {
		return nil, fmt.Errorf("GetFollowedAndFollowers: %v", err)
	}
	defer rows.Close()

	var followedAndFollowers []User
	for rows.Next() {
		var userId int
		if err := rows.Scan(&userId); err != nil {
			return nil, fmt.Errorf("GetFollowedAndFollowers scan: %v", err)
		}
		user := User{Online: false}
		if err := user.GetUserById(db, userId); err != nil {
			return nil, fmt.Errorf("GetFollowedAndFollowers GetUserById: %v", err)
		}
		followedAndFollowers = append(followedAndFollowers, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetFollowedAndFollowers rows.Err: %v", err)
	}

	return followedAndFollowers, nil
}

func (u *User) GetFollowedAndFollowersNotInGroup(db *sql.DB, groupId int) ([]User, error) {
	query := `
        SELECT id FROM (
            SELECT User_id AS id FROM Follow WHERE Follower_id = 9
            UNION
            SELECT Follower_id AS id FROM Follow WHERE User_id = 9
        ) AS combined
        WHERE id NOT IN (SELECT User_id FROM joinner WHERE Group_id = 6)
    `
	rows, err := db.Query(query, u.ID, u.ID, groupId)
	if err != nil {
		return nil, fmt.Errorf("GetFollowedAndFollowersNotInGroup: %v", err)
	}
	defer rows.Close()

	var followedAndFollowers []User
	for rows.Next() {
		var userId int
		if err := rows.Scan(&userId); err != nil {
			return nil, fmt.Errorf("GetFollowedAndFollowersNotInGroup scan: %v", err)
		}
		user := User{}
		if err := user.GetUserById(db, userId); err != nil {
			return nil, fmt.Errorf("GetFollowedAndFollowersNotInGroup GetUserById: %v", err)
		}
		followedAndFollowers = append(followedAndFollowers, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetFollowedAndFollowersNotInGroup rows.Err: %v", err)
	}

	return followedAndFollowers, nil
}

func (u *User) GetFollowedAndFollowers(db *sql.DB) ([]User, error) {
	query := `
        SELECT id FROM 
            SELECT User_id AS id FROM Follow WHERE Follower_id = ?
            UNION
            SELECT Follower_id AS id FROM Follow WHERE User_id = ?
    `
	rows, err := db.Query(query, u.ID, u.ID)
	if err != nil {
		return nil, fmt.Errorf("GetFollowedAndFollowersNotInGroup: %v", err)
	}
	defer rows.Close()

	var followedAndFollowers []User
	for rows.Next() {
		var userId int
		if err := rows.Scan(&userId); err != nil {
			return nil, fmt.Errorf("GetFollowedAndFollowersNotInGroup scan: %v", err)
		}
		user := User{}
		if err := user.GetUserById(db, userId); err != nil {
			return nil, fmt.Errorf("GetFollowedAndFollowersNotInGroup GetUserById: %v", err)
		}
		followedAndFollowers = append(followedAndFollowers, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetFollowedAndFollowersNotInGroup rows.Err: %v", err)
	}

	return followedAndFollowers, nil
}
