package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"social_network/models"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {

	NewPost := models.Post{}
	
	err := json.NewDecoder(r.Body).Decode(&NewPost)

	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		http.Error(w, "400 bad request.", http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"Error": err})
		return
	}
	
	NewPost.Check()

	idPost, creationError := NewPost.Create(DB)

	if creationError != nil {
		fmt.Println("Error creating post :", creationError)
		http.Error(w, "400 bad request.", http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"Error": creationError})
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"success": idPost})
}

