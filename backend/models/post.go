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
	Privacy      string    `json:"privacy"`
	PostGroup    string    `json:"postGroup"`
	CreationDate time.Time `json:"creationDate"`
	AllowedUsers []int
}

//

type FeedPost struct {
	Post
	GroupName   string `json:"groupName"`
	Description string `json:"description"`
}

type Comment struct {
	IdComment int `json:"id"`
	FirtsName string
	LastName  string
	Content   string
	Type      string
}

type PostDetails struct {
	Post
	GroupName   string `json:"groupName"`
	Description string `json:"description"`
	Comments    []Comment
}

// func (P *Post) GetAll(controllerDB *sql.DB) ([]Post,error) {
// }

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

	if P.Privacy == "almostprivate" {
		er := RegisterAllowedUsers(P.AllowedUsers, idPost, controllerDB)
		if er != nil {
			return -1, er
		}
	}
	return idPost, nil
}

func RegisterAllowedUsers(users []int, idPost int, controllerDB *sql.DB) error {
	statement, err := controllerDB.Prepare("INSERT INTO AllowedPost (Post_id , User_id) VALUES (?,?)")
	if err != nil {
		return err
	}
	for _, User_id := range users {
		_, er := statement.Exec(idPost, User_id)
		if er != nil {
			return er
		}
	}
	return nil
}

func (P *Post) Check() bool {
	return P.User_id != 0 && P.Content != "" && P.Titre != "" && P.CheckPrivacy()
}

func (P *Post) CheckPrivacy() bool {
	return P.Privacy == "public" || P.Privacy == "private" || P.Privacy == "almostprivate" || P.Privacy == "groupe"
}

func (P *Post) CheckGroupMember(DB *sql.DB) (int, error) {
	statement, err := DB.Prepare("select COUNT(id) from Joinner where User_id = ? and Group_id = ? ")
	if err != nil {
		return 0, err
	}
	var num int
	statement.QueryRow(P.User_id, P.Group_id).Scan(&num)
	return num, nil
}

func (P *PostDetails) GetPost(DB *sql.DB ,UserID, post_id int ) error {
	statement, err := DB.Prepare(`
	SELECT P.id , P.Group_id , P.titre , P.image , P.content , P.privacy , P.creationDate , G.titre , G.description
	FROM Post as P 
   	LEFT JOIN "Group" as G on P.Group_id = G.id
   	WHERE P.id = ? and ( P.privacy = "public" or (P.privacy = "private" and P.User_id in 
	   (SELECT F.User_id from Follow F WHERE F.Follower_id = ? )) or
	   		P.privacy = "almostprivate" and P.id in 
		   (SELECT A.Post_id  from AllowedPost as A WHERE A.User_id = ? ) or
			P.Group_id in 
			   (SELECT J.Group_id from Joinner as J WHERE J.User_id = ?)) 
	`)
	if err != nil {
		return err
	}
	var sqler error
	sqler = statement.QueryRow( post_id ,UserID, UserID , UserID ).Scan(&P.Post.Id, &P.Post.Group_id, &P.Post.Titre, &P.Post.Image, &P.Post.Content, &P.Post.Privacy, &P.Post.CreationDate, &P.GroupName, &P.Description )

	if sqler == sql.ErrNoRows {
		return sqler
	}	
	return nil
}

func (P *PostDetails) Getcomments(DB *sql.DB) error {
	statement, err := DB.Prepare(`
	select C.id , U.firstName , U.lastName , C.comment , C.type from Comment as C 
JOIN User as U on U.id = C.User_id
where C.Post_id = ?
	`)
	if err != nil {
		return err
	}
	lines, er := statement.Query(P.Post.Id)
	if er != nil {
		return er
	}
	for lines.Next() {
		com := Comment{}
		lines.Scan(&com.IdComment, &com.FirtsName, &com.LastName, &com.Content, &com.Type)
		P.Comments = append(P.Comments, com)
	}
	return nil
}
