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

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	NewPost := models.Post{User_id: 2}

	err := json.NewDecoder(r.Body).Decode(&NewPost)

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

	user := models.User{}

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
	UserID := 2

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

	data, err := json.Marshal(post)
	if err != nil {
		helper.ErrorPage(w, 500)
		return
	}
	w.Write(data)
}

func GroupPost(w http.ResponseWriter, r *http.Request) {
	// variable temporaire
	UserID := 2
	group_id, er := strconv.Atoi(r.URL.Query().Get("groupid"))
	if er != nil {
		helper.ErrorPage(w, 400)
		return
	}
	group := models.GroupeInfo{
		Id: group_id,
	}
	ismenber, Er := group.IsMember(DB, UserID)

	// internal server error
	if Er != nil && Er != sql.ErrNoRows {
		helper.ErrorPage(w, 500)
		return
	}

	// you're not a member of this group
	if !ismenber {
		helper.ErrorPage(w, 400)
		return
	}
	PostEr := group.GetGroupPost(DB, UserID)
	if PostEr != nil {
		helper.ErrorPage(w, 500)
		return
	}
	data, err := json.Marshal(group)
	if err != nil {
		helper.ErrorPage(w, 500)
	}
	w.Write(data)
}
