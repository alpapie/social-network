package route

import (
	"net/http"
	"social_network/controller"
	"social_network/middleware"
)

func Routes() {
	http.HandleFunc("/", middleware.CheckMethod(controller.Login, "get"))
	http.HandleFunc("/server/post", middleware.IsAuth(middleware.CheckMethod(controller.PostHandler, "post")))
	http.HandleFunc("/server/getposts", middleware.IsAuth(middleware.CheckMethod(controller.GetPost, "get")))
	http.HandleFunc("/server/getpost", middleware.IsAuth(middleware.CheckMethod(controller.PostDetail, "get")))

}
