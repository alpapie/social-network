package route

import (
	"net/http"
	"social_network/controller"
	"social_network/middleware"
)

func Routes() {
	http.HandleFunc("/", middleware.CheckMethod(controller.Home, "get"))

	http.HandleFunc("/server/login", middleware.Ispath(middleware.CheckMethod(controller.Login, "post"), "login"))
	http.HandleFunc("/server/register", controller.RegisterUser)
}
