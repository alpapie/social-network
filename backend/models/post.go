package models

import (
	"database/sql"
	"strings"
	"time"
)

type Post struct {
	Id           int    `json:"id"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Group_id     int    `json:"Group_id"`
	User_id      int    `json:"User_id" `
	Titre        string `json:"titre"`
	Image        string `json:"image"`
	Content      string `json:"content"`
	Privacy      string `json:"privacy"`
	PostGroup    string `json:"postGroup"`
	CreationDate string `json:"creationDate"`
	AllowedUsers []int
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
type GroupeInfo struct {
	Id          int
	Title       string
	Description string
	Posts       []Post
}

func (P *PostDetails) GetPost(DB *sql.DB, UserID, post_id int) error {
	statement, err := DB.Prepare(`
	SELECT U.id , U.firstName , U.lastName , P.id , coalesce(P.Group_id,0) as Group_id , P.titre ,  coalesce(P.image, '') as image , P.content , P.privacy , coalesce(P.creationDate,'') as creationDate , coalesce(G.titre,'') as titre , coalesce(G.description,'') as description
	FROM Post as P 
	JOIN User as U on U.id = P.User_id
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
	sqlErr := statement.QueryRow(post_id, UserID, UserID, UserID).Scan(&P.Post.User_id, &P.Post.FirstName, &P.Post.LastName, &P.Post.Id, &P.Post.Group_id, &P.Post.Titre, &P.Post.Image, &P.Post.Content, &P.Post.Privacy, &P.Post.CreationDate, &P.GroupName, &P.Description)
	if sqlErr == sql.ErrNoRows {
		return sqlErr
	}
	return nil
}

func (P *PostDetails) GetComments(DB *sql.DB) error {
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
		lines.Scan(&com.Id, &com.FirstName, &com.LastName, &com.Comment, &com.Type)
		P.Comments = append(P.Comments, com)
	}
	return nil
}

func (U *User) GetPosts(controllerDB *sql.DB) ([]FeedPost, error) {
	// variable temporaire
	UserID := 2

	statement, err := controllerDB.Prepare(`
	SELECT U.id , U.firstName , U.lastName , P.id , coalesce(P.Group_id,0) as Group_id , P.titre , coalesce(P.image ,'') as image , P.content , P.privacy , coalesce(P.creationDate,"") as creationDate ,coalesce(G.titre,"") as groupName ,coalesce(G.description,"") as description
	FROM Post as P 
	JOIN User as U on U.id = P.User_id
   	LEFT JOIN "Group" as G on P.Group_id = G.id
   	WHERE P.privacy = "public" or (P.privacy = "private" and P.User_id in 
	   (SELECT F.User_id from Follow as F WHERE F.Follower_id = ? )) or
	   		P.privacy = "almostprivate" and P.id in 
		   (SELECT A.Post_id  from AllowedPost as A WHERE A.User_id = ? ) or
			P.Group_id in 
			   (SELECT J.Group_id from Joinner as J WHERE J.User_id = ?) LIMIT 10 OFFSET ?;
	`)
	if err != nil {
		return []FeedPost{}, err
	}

	posts := []FeedPost{}

	lines, err := statement.Query(UserID, UserID, UserID, 0)
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

	statement, err := controllerDB.Prepare("INSERT INTO Post (Group_id, User_id, titre, image, content, privacy, postGroup,creationDate) VALUES (?,?,?,?,?,?,?,?)")
	if err != nil {
		return -1, err
	}

	sqlResult, err := statement.Exec(P.Group_id, P.User_id, P.Titre, P.Image, P.Content, P.Privacy, P.PostGroup, time.Now().String())
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
	return P.User_id != 0 && strings.TrimSpace(P.Content) != "" && strings.TrimSpace(P.Titre) != "" && P.CheckPrivacy()
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

func (G *GroupeInfo) IsMember(DB *sql.DB, UserId int) (bool, error) {
	statement, err := DB.Prepare(`
	SELECT  G.titre , G.description from Joinner as J JOIN "Group" as G 
	on G.id = J.Group_id
	WHERE J.Group_id = ? and J.User_id = ?
	`)
	if err != nil {
		return false, err
	}
	Er := statement.QueryRow(G.Id, UserId).Scan(&G.Title, &G.Description)
	if Er != nil {
		return false, Er
	}
	return true, nil
}

func (G *GroupeInfo) GetGroupPost(DB *sql.DB, userId int) error {
	statement, err := DB.Prepare(`
	SELECT 
	U.id , U.firstName , U.lastName , P.id , coalesce(P.Group_id , 0) as Group_id , P.titre ,coalesce(P.image, '') as image , P.content , P.privacy , coalesce(P.creationDate , '') as creationDate
	from Post as P
	JOIN User as U on U.id = P.User_id
	where P.Group_id = ?
	`)
	if err != nil {
		return err
	}
	posts, Er := statement.Query(G.Id)
	if Er != nil && Er != sql.ErrNoRows {
		return Er
	}
	for posts.Next() {
		p := Post{}
		posts.Scan(&p.User_id ,&p.FirstName , &p.LastName ,&p.Id, &p.Group_id, &p.Titre, &p.Image, &p.Content, &p.Privacy, &p.CreationDate)
		G.Posts = append(G.Posts, p)
	}
	return nil
}