package route

import (
	"net/http"
	"social_network/controller"
	"social_network/middleware"
)

func Routes() {
	http.HandleFunc("/", middleware.CheckMethod(controller.Login, "get"))
	http.HandleFunc("/server/addPost", middleware.IsAuth(middleware.CheckMethod(controller.CreatePostHandler, "post")))
	http.HandleFunc("/server/getPost", middleware.IsAuth(middleware.CheckMethod(controller.PostDetail, "get")))
	http.HandleFunc("/server/getPosts", middleware.IsAuth(middleware.CheckMethod(controller.PostsByUserHandler, "get")))
	http.HandleFunc("/server/group", middleware.IsAuth(middleware.CheckMethod(controller.GroupPost, "get")))
	http.HandleFunc("/server/addComment", middleware.IsAuth(middleware.CheckMethod(controller.AddCommentHandler, "post")))
}
