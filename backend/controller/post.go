package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	helper "social_network/helper"
	"social_network/models"
	"strconv"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
}

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, _, UserID := helper.Auth(DB, r)
	NewPost := models.Post{User_id: UserID}

	err := json.NewDecoder(r.Body).Decode(&NewPost)
	fmt.Println("here is the decoded post ", NewPost)
	if err != nil || !NewPost.Check() {
		fmt.Println("Error decoding JSON:", err)
		http.Error(w, "400 bad request.", http.StatusBadRequest)
		return
	}

	// if the Post belong to a group Check first if the user is member of the group
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
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{"success": idPost})
}

func PostsByUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, _, UserID := helper.Auth(DB, r)
	user := models.User{ID: UserID}

	posts, err := user.GetPosts(DB)
	if err != nil {
		helper.ErrorPage(w, 500)
		return
	}

	err = json.NewEncoder(w).Encode(posts)
	if err != nil {
		helper.ErrorPage(w, 500)
		return
	}
}

func PostDetail(w http.ResponseWriter, r *http.Request) {

	post_id, er := strconv.Atoi(r.URL.Query().Get("postid"))
	if er != nil {
		helper.ErrorPage(w, 400)
		return
	}
	// variable temporaire
	_, _, UserID := helper.Auth(DB, r)

	post := models.PostDetails{}
	Er := post.GetPost(DB, UserID, post_id)

	if Er != nil {
		// You don't have acces to this post or the post does not exist
		if Er == sql.ErrNoRows {
			helper.ErrorPage(w, 400)
			return
		}
		// Internal server error
		helper.ErrorPage(w, 500)
		return
	}

	// Get comments of the post
	ComErr := post.GetComments(DB)
	if ComErr != nil {
		helper.ErrorPage(w, 500)
		return
	}
	helper.WriteJSON(w,200,map[string]interface{}{"data":post},r.Header)
}

