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
	http.HandleFunc("/server/groups", middleware.Ispath(middleware.CheckMethod(controller.GetAllNotjoinedGroups, "get"), "groups"))
	http.HandleFunc("/server/createGroup", controller.CreateGroupHandler)
	http.HandleFunc("/server/invitegroup", middleware.Ispath(middleware.CheckMethod(controller.CreateInvitationGroup, "post"), "invitegroup"))
	http.HandleFunc("/server/followgroup", middleware.Ispath(middleware.CheckMethod(controller.CreateFollowGroup, "post"), "followgroup"))
	http.HandleFunc("/server/createEvent", middleware.Ispath(middleware.CheckMethod(controller.CreateEvent, "post"), "createEvent"))
	http.HandleFunc("/server/getgroupdetail/", controller.GetGroupDetail)
}
