package models

import (
	"database/sql"
	"fmt"
)

func (u *User) GetUnFollow(DB *sql.DB, limit int) ([]User, error) {
	req := `SELECT id,"firstName","lastName", email,avatar 
	FROM user
	WHERE id NOT IN (
		SELECT user_id
		FROM follow
		WHERE follower_id = ?
	)
	and id!=? LIMIT ? `
	row, err := DB.Query(req, u.ID, u.ID, limit)
	users := []User{}
	if err != nil {
		return users, err
	}
	users = ExtractUserData(row, users)

	return users, nil
}

func (U *User) Folower(DB *sql.DB, user_id int) ([]User, error) {
	req := `SELECT  u.id,u."firstName",u."lastName", u.email,u.avatar from "User" AS "u" inner JOIN "Follow" AS "f" on f."Follower_id"=u.id  where u.id!=?`
	row, err := DB.Query(req, user_id)
	users := []User{}
	if err != nil {
		return users, err
	}
	return ExtractUserData(row, users), nil
}

func (U *User) Following(DB *sql.DB, user_id int) ([]User, error) {
	req := ` SELECT u.id,u."firstName",u."lastName", u.email,u.avatar from "User" AS "u" inner JOIN "Follow" AS "f" on f."User_id"=u.id  where u.id!=?`
	row, err := DB.Query(req, user_id)

	users := []User{}
	if err != nil {
		return users, err
	}
	return ExtractUserData(row, users), nil
}

func ExtractUserData(row *sql.Rows, users []User) []User {
	newUser:= User{}
	for row.Next() {
		row.Scan(&newUser.ID, &newUser.FirstName, &newUser.LastName, &newUser.Email, &newUser.Avatar)
		users = append(users, newUser)
	}
	return users
}

func (U *User) CreatedPost(DB *sql.DB, user_id int) ([]FeedPost, error) {
	req := `SELECT p.id,p.titre,p.image,p.content,p.privacy ,p."Group_id",g.title,u."firstName",u."lastName",u.avatar,p."creationDate" FROM "Post" "p" LEFT JOIN "Group" "g" ON p."Group_id"=g.id  join user "u" on p."User_id"=u.id  where p."User_id"=? `
	row, err := DB.Query(req, user_id)
	posts := []FeedPost{}
	if err != nil {
		return posts, err
	}
	for row.Next() {
		post := FeedPost{}
		row.Scan(&post.Id, &post.Titre, &post.Image, &post.Content, &post.Privacy, &post.Group_id, &post.GroupName, &post.FirstName, &post.LastName, &post.Avatar, &post.CreationDate)
		fmt.Println(len(posts),"created postin the models")
		posts = append(posts, post)
	}
	return posts, nil
}

func (U *User) UpdateStatus(DB *sql.DB) error {
	req := `UPDATE User SET ispublic = ? WHERE id=?`

	_, err := DB.Exec(req, U.IsPublic,U.ID)
	return err
}

func (U *User)GetFollowerAndFollowing(DB *sql.DB,user_id int) ([]User, error ){
	req:=`SELECT  u.id,u."firstName",u."lastName", u.email,u.avatar from "User" AS "u" inner JOIN "Follow" AS "f" on f."Follower_id"=u.id or f."User_id"=u.id where u.id!=?`
	row, err := DB.Query(req, user_id)
	users := []User{}
	if err != nil {
		return users, err
	}
	return ExtractUserData(row, users), nil
}