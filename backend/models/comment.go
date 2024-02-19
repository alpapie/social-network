package models

import (
	"database/sql"
	"html"
	"strings"
)

type Comment struct {
	Id           int    `json:"id"`
	Comment      string `json:"comment"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	UserId       int    `json:"userId"`
	PostId       int    `json:"postId"`
	CreationDate string `json:"creationDate"`
	Image        string `json:"image"`
}

func (C *Comment) Create(DB *sql.DB) (int, error) {
	statement, err := DB.Prepare("Insert INTO Comment (post_id, user_id, comment, image) values (?,?,?,?)")
	if err != nil {
		return -1, err
	}
	defer statement.Close()

	sqlResult, err := statement.Exec(C.PostId, C.UserId, C.Comment,C.Image)
	if err != nil {
		return -1, err
	}

	lastInsertedId, err := sqlResult.LastInsertId()
	if err != nil {
		return -1, err
	}

	return int(lastInsertedId), nil
}

func (C *Comment) IsAllowedToComment(DB *sql.DB) (int, error) {
	stmt, err := DB.Prepare(`SELECT count(P.id) 
	FROM Post as P
   	LEFT JOIN "Group" as G on P.Group_id = G.id
   	WHERE P.id = ? and ( P.privacy = "public" or (P.privacy = "private" and P.User_id in 
	   (SELECT F.User_id from Follow F WHERE F.Follower_id = ? )) or
	   		P.privacy = "almostprivate" and P.id in 
		   (SELECT A.Post_id  from AllowedPost as A WHERE A.User_id = ? ) or
			P.Group_id in 
			   (SELECT J.Group_id from Joinner as J WHERE J.User_id = ?)) `)

	if err != nil {
		return 0, err
	}

	resultCount := 0
	row := stmt.QueryRow(C.PostId, C.UserId, C.UserId, C.UserId)
	err = row.Scan(&resultCount)
	if err != nil {
		return 0, err
	}
	return resultCount, nil
}

func (c *Comment) ValidateComment() bool {
	c.Comment = html.EscapeString(strings.TrimSpace(c.Comment))
	return c.Comment != "" && c.PostId != 0
}
