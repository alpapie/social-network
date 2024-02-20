package models

import (
	"database/sql"
	"fmt"
	"html"
	"strings"
	"time"
)

type Post struct {
	Id           int    `json:"id"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Group_id     int    `json:"Group_id"`
	User_id      int    `json:"User_id"`
	Titre        string `json:"titre"`
	Image        string `json:"image"`
	Content      string `json:"content"`
	Privacy      string `json:"privacy"`
	CreationDate string `json:"creationDate"`
	Avatar	string	`json:"groupe_title"`
	// AllowedUsers []int
	AllowedUsers []int  `json:allowedusers`
}

type FeedPost struct {
	Post
	GroupName   string `json:"groupName"`
	Description string `json:"description"`
}

type PostDetails struct {
	Post
	GroupName   string `json:"groupName"`
	Description string `json:"description"`
	Comments    []Comment
}


func (P *PostDetails) GetPost(DB *sql.DB, UserID, post_id int) error {
	statement, err := DB.Prepare(`
	SELECT U.id , U.firstName , U.lastName , P.id , coalesce(P.Group_id,0) as Group_id , P.titre ,  coalesce(P.image, '') as image , P.content , P.privacy , coalesce(P.creationDate,'') as creationDate , coalesce(G.title,'') as title , coalesce(G.description,'') as description
	FROM Post as P 
	JOIN User as U on U.id = P.User_id
   	LEFT JOIN "Group" as G on P.Group_id = G.id
   	WHERE P.id = ? and ( P.User_id = ? OR P.privacy = "public" or (P.privacy = "private" and P.User_id in 
	   (SELECT F.User_id from Follow F WHERE F.Follower_id = ? )) or
	   		P.privacy = "almostprivate" and P.id in 
		   (SELECT A.Post_id  from AllowedPost as A WHERE A.User_id = ? ) or
			P.Group_id in 
			   (SELECT J.Group_id from Joinner as J WHERE J.User_id = ?)) 
	`)
	if err != nil {
		return err
	}
	sqlErr := statement.QueryRow(post_id, UserID, UserID, UserID, UserID).Scan(&P.Post.User_id, &P.Post.FirstName, &P.Post.LastName, &P.Post.Id, &P.Post.Group_id, &P.Post.Titre, &P.Post.Image, &P.Post.Content, &P.Post.Privacy, &P.Post.CreationDate, &P.GroupName, &P.Description)
	if sqlErr == sql.ErrNoRows {
		return sqlErr
	}
	return nil
}

func (P *PostDetails) GetComments(DB *sql.DB) error {
	statement, err := DB.Prepare(`
	select C.id , U.firstName , U.lastName , C.comment ,C.creationDate, coalesce(C.image, '') as image ,  C.User_id, C.Post_id from Comment as C 
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
		lines.Scan(&com.Id, &com.FirstName, &com.LastName, &com.Comment, &com.CreationDate,&com.Image, &com.UserId, &com.PostId)
		P.Comments = append(P.Comments, com)
	}
	return nil
}

func (U *User) GetPosts(controllerDB *sql.DB) ([]FeedPost, error) {

	statement, err := controllerDB.Prepare(`
	SELECT U.id , U.firstName , U.lastName , P.id , coalesce(P.Group_id,0) as Group_id , P.titre , coalesce(P.image ,'') as image , P.content , P.privacy , coalesce(P.creationDate,"") as creationDate ,coalesce(G.title,"") as groupName ,coalesce(G.description,"") as description
	FROM Post as P 
	JOIN User as U on U.id = P.User_id
   	LEFT JOIN "Group" as G on P.Group_id = G.id
   	WHERE P.privacy = "public" or 
	P.User_id = ?
	or (P.privacy = "private" and P.User_id in 
	   (SELECT F.User_id from Follow as F WHERE F.Follower_id = ? )) or
	   		P.privacy = "almostprivate" and P.id in 
		   (SELECT A.Post_id  from AllowedPost as A WHERE A.User_id = ? ) or
			P.Group_id in 
			   (SELECT J.Group_id from Joinner as J WHERE J.User_id = ?) order by P.id desc;
	`)
	if err != nil {
		return []FeedPost{}, err
	}

	fmt.Println("the user id", U.ID)

	posts := []FeedPost{}

	lines, err := statement.Query(U.ID, U.ID, U.ID, U.ID)
	if err != nil {
		return []FeedPost{}, err
	}
	defer lines.Close()

	for lines.Next() {
		post := FeedPost{}
		err = lines.Scan(&post.User_id, &post.FirstName, &post.LastName, &post.Id, &post.Group_id, &post.Titre, &post.Image, &post.Content, &post.Privacy, &post.CreationDate, &post.GroupName, &post.Description)
		if err != nil {
			return []FeedPost{}, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (P *Post) Create(controllerDB *sql.DB) (int, error) {

	statement, err := controllerDB.Prepare("INSERT INTO Post (Group_id, User_id, titre, image, content, privacy, creationDate) VALUES (?,?,?,?,?,?,?)")
	if err != nil {
		return -1, err
	}

	sqlResult, err := statement.Exec(P.Group_id, P.User_id, P.Titre, P.Image, P.Content, P.Privacy, time.Now().String())
	if err != nil {
		return -1, err
	}

	lastInsertedId, err := sqlResult.LastInsertId()
	if err != nil {
		return -1, err
	}
	var idPost = int(lastInsertedId)

	// Add the post creator to allowed users
	P.AllowedUsers = append(P.AllowedUsers, P.User_id)

	if P.Group_id != 0 {
		P.Privacy = "groupe"
	}

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
	P.Content = html.EscapeString(strings.TrimSpace(P.Content))
	P.Titre = html.EscapeString(strings.TrimSpace(P.Titre))

	return P.User_id != 0 && P.Content != "" && P.CheckPrivacy()
}

func (P *Post) CheckPrivacy() bool {
	privacy := strings.TrimSpace(strings.ToLower(P.Privacy))
	return privacy == "public" || privacy == "private" || privacy == "almostprivate" || privacy == "groupe"
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


func GetGroupPost(DB *sql.DB, groupId int) ([]Post ,error) {
	posts := []Post{}
	statement, err := DB.Prepare(`
	SELECT 
	U.id , U.firstName , U.lastName , P.id , coalesce(P.Group_id , 0) as Group_id , P.titre ,coalesce(P.image, '') as image , P.content , P.privacy , coalesce(P.creationDate , '') as creationDate
	from Post as P
	JOIN User as U on U.id = P.User_id
	where P.Group_id = ? order by P.id desc;
	`)
	if err != nil {
		return posts , err
	}
	postsdata, Er := statement.Query(groupId)
	if Er != nil && Er != sql.ErrNoRows {
		return posts, Er
	}
	for postsdata.Next() {
		p := Post{}
		postsdata.Scan(&p.User_id, &p.FirstName, &p.LastName, &p.Id, &p.Group_id, &p.Titre, &p.Image, &p.Content, &p.Privacy, &p.CreationDate)
		posts = append(posts, p)
	}
	return posts ,nil
}

