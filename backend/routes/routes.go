package route

import (
	"net/http"
	"social_network/controller"
	"social_network/middleware"
)

func Routes() {
	http.HandleFunc("/server/", middleware.Ispath(middleware.CheckMethod(controller.Auth, "get"),""))
	http.HandleFunc("/server/home", middleware.Log(middleware.Ispath(middleware.CheckMethod(controller.Home, "get"),"home")))

	http.HandleFunc("/server/login", middleware.Ispath(middleware.CheckMethod(controller.Login, "post"), "login"))
	http.HandleFunc("/server/profile", middleware.Ispath(middleware.CheckMethod(controller.Profil, "get"), "profile"))
	http.HandleFunc("/server/unfollower", middleware.Ispath(middleware.CheckMethod(controller.UnFollow, "get"), "unfollower"))
	http.HandleFunc("/server/follow", middleware.Log(middleware.CheckMethod(controller.Follow, "get")))
	http.HandleFunc("/server/changestatus", middleware.Log(middleware.Ispath(middleware.CheckMethod(controller.UpdateProfil, "get"),"changestatus")))

	http.HandleFunc("/server/logout", middleware.Log(middleware.Ispath(middleware.CheckMethod(controller.Logout, "get"), "logout")))

	http.HandleFunc("/server/register", middleware.Ispath(middleware.CheckMethod(controller.RegisterUser,"post"),"register"))

	http.HandleFunc("/server/addPost", middleware.Log(middleware.CheckMethod(controller.CreatePostHandler, "post")))
	http.HandleFunc("/server/getPost", middleware.Log(middleware.CheckMethod(controller.PostDetail, "get")))
	http.HandleFunc("/server/getPosts", middleware.Log(middleware.CheckMethod(controller.PostsByUserHandler, "get")))
	http.HandleFunc("/server/group", middleware.Log(middleware.CheckMethod(controller.GroupPost, "get")))
	http.HandleFunc("/server/addComment", middleware.Log(middleware.CheckMethod(controller.AddCommentHandler, "post")))
}
