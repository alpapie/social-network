package route

import (
	"net/http"
	"social_network/controller"
	"social_network/middleware"
)

func Routes() {
	http.HandleFunc("/server/", middleware.Ispath(middleware.CheckMethod(controller.Auth, "get"), ""))
	http.HandleFunc("/server/home", middleware.Log(middleware.Ispath(middleware.CheckMethod(controller.Home, "get"), "home")))

	http.HandleFunc("/server/login", middleware.Ispath(middleware.CheckMethod(controller.Login, "post"), "login"))
	http.HandleFunc("/server/profile", middleware.Log(middleware.Ispath(middleware.CheckMethod(controller.Profil, "get"), "profile")))
	http.HandleFunc("/server/unfollower", middleware.Log(middleware.Ispath(middleware.CheckMethod(controller.UnFollow, "get"), "unfollower")))
	http.HandleFunc("/server/follow", middleware.Log(middleware.Ispath(middleware.CheckMethod(controller.Follow, "get"), "follow")))
	http.HandleFunc("/server/changestatus", middleware.Log(middleware.Ispath(middleware.CheckMethod(controller.UpdateProfil, "get"), "changestatus")))

	http.HandleFunc("/server/logout", middleware.Log(middleware.Ispath(middleware.CheckMethod(controller.Logout, "get"), "logout")))

	http.HandleFunc("/server/register", middleware.Ispath(middleware.CheckMethod(controller.RegisterUser, "post"), "register"))

	http.HandleFunc("/server/addPost", middleware.Log(middleware.CheckMethod(controller.CreatePostHandler, "post")))
	http.HandleFunc("/server/getPost", middleware.Log(middleware.CheckMethod(controller.PostDetail, "get")))
	http.HandleFunc("/server/getPosts", middleware.Log(middleware.CheckMethod(controller.PostsByUserHandler, "get")))
	http.HandleFunc("/server/addComment", middleware.Log(middleware.CheckMethod(controller.AddCommentHandler, "post")))
	// http.HandleFunc("/server/", middleware.Ispath(middleware.CheckMethod(controller.Auth, "get"), ""))

	http.HandleFunc("/server/groups", middleware.Log(middleware.Ispath(middleware.CheckMethod(controller.GetAllNotjoinedGroups, "get"), "groups")))
	http.HandleFunc("/server/createGroup", middleware.Log(controller.CreateGroupHandler))
	http.HandleFunc("/server/invitegroup", middleware.Log(middleware.Ispath(middleware.CheckMethod(controller.CreateInvitationGroup, "post"), "invitegroup")))
	http.HandleFunc("/server/followgroup", middleware.Log(middleware.Ispath(middleware.CheckMethod(controller.CreateFollowGroup, "post"), "followgroup")))
	http.HandleFunc("/server/createEvent", middleware.Log(middleware.Ispath(middleware.CheckMethod(controller.CreateEvent, "post"), "createEvent")))
	http.HandleFunc("/server/getoption", middleware.Log(middleware.CheckMethod(controller.GetOption, "get")))
	http.HandleFunc("/server/getgroupdetail/", middleware.Log(controller.GetGroupDetail))

	http.HandleFunc("/server/chat", middleware.Log(middleware.Ispath(middleware.CheckMethod(controller.GetChatMessageHandler, "get"), "chat")))
	http.HandleFunc("/server/chatgroup", middleware.Log(middleware.Ispath(middleware.CheckMethod(controller.GetGroupMessageHandler, "get"), "chatgroup")))

	//socket Handlers
	http.HandleFunc("/server/ws", middleware.Log(middleware.Ispath(middleware.CheckMethod(controller.WsHandler, "get"), "ws")))

}
