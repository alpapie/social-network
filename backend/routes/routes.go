package route

import (
	"net/http"
	"social_network/controller"
	"social_network/middleware"
)

func Routes() {
	http.HandleFunc("/server/", middleware.Ispath(middleware.CheckMethod(controller.Auth, "get"), ""))

	http.HandleFunc("/server/login", middleware.Ispath(middleware.CheckMethod(controller.Login, "post"), "login"))
	http.HandleFunc("/server/register", middleware.Ispath(middleware.CheckMethod(controller.RegisterUser, "post"), "register"))

	// http.HandleFunc("/", middleware.CheckMethod(controller.Login, "get"))
	http.HandleFunc("/server/home", middleware.IsAuth(middleware.CheckMethod(controller.Home, "get")))
	http.HandleFunc("/server/addPost", middleware.IsAuth(middleware.CheckMethod(controller.CreatePostHandler, "post")))
	http.HandleFunc("/server/getPost", middleware.IsAuth(middleware.CheckMethod(controller.PostDetail, "get")))
	http.HandleFunc("/server/getPosts", middleware.IsAuth(middleware.CheckMethod(controller.PostsByUserHandler, "get")))
	http.HandleFunc("/server/group", middleware.Log(middleware.CheckMethod(controller.GroupPost, "get")))
	http.HandleFunc("/server/addComment", middleware.IsAuth(middleware.CheckMethod(controller.AddCommentHandler, "post")))
}
