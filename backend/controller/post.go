package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"social_network/models"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	NewPost := models.Post{}

	err := json.NewDecoder(r.Body).Decode(&NewPost)

	if err != nil || !NewPost.Check() {
		fmt.Println("Error decoding JSON:", err)
		http.Error(w, "400 bad request.", http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"Error": err})
		return
	}
	// if the Post belong to a group Check fist if the user is member of the group
	if NewPost.Group_id != 0 {
		num, Err := NewPost.CheckGroupMember(DB)
		if Err != nil {
			fmt.Println("Error server ", Err)
			http.Error(w, "500 bad request.", http.StatusInternalServerError)
			return
		}
		if num != 1 {
			fmt.Println("you're not a member of the group")
			http.Error(w, "400 bad request.", http.StatusBadRequest)
			return
		}
	}
	idPost, creationError := NewPost.Create(DB)

	if creationError != nil {
		fmt.Println("Error creating post :", creationError)
		http.Error(w, "400 bad request.", http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"Error": creationError})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{"success": idPost})
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	//variable temporaire
	UserID := 2
	statement, err := DB.Prepare(`
	SELECT P.id , P.Group_id , P.titre , P.image , P.content , P.privacy , P.creationDate , G.titre , G.description
	FROM Post as P 
   	LEFT JOIN "Group" as G on P.Group_id = G.id
   	WHERE P.privacy = "public" or (P.privacy = "private" and P.User_id in 
	   (SELECT F.User_id from Follow F WHERE F.Follower_id = ? )) or
	   		P.privacy = "almostprivate" and P.id in 
		   (SELECT A.Post_id  from AllowedPost as A WHERE A.User_id = ? ) or
			P.Group_id in 
			   (SELECT J.Group_id from Joinner as J WHERE J.User_id = ?) LIMIT 10 OFFSET ?
			
	`)
	if err != nil {
		Error(w, err, 500)
		return
	}

	posts := []models.FeedPost{}

	lines, er := statement.Query(UserID, UserID, UserID, 0)
	if err != nil {
		Error(w, er, 500)
		return
	}
	for lines.Next() {
		post := models.FeedPost{}
		lines.Scan(&post.Post.Id, &post.Post.Group_id, &post.Post.Titre, &post.Post.Image, &post.Post.Content, &post.Post.Privacy, &post.Post.CreationDate, &post.GroupName, &post.Description)
		posts = append(posts, post)
	}
	data, err := json.Marshal(posts)
	if err != nil {
		Error(w, err, 500)
	}
	w.Write(data)
	fmt.Println("here are the users", string(data))
}

func Error(w http.ResponseWriter, er error, statusCode int) {
	fmt.Println("Error ", statusCode, er)
	http.Error(w, "Error ", statusCode)
	return
}
