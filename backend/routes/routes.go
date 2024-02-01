package route

import (
	"net/http"
	"social_network/controller"
	"social_network/middleware"
)

func Routes() {
	http.HandleFunc("/", middleware.CheckMethod(controller.Login, "get"))
	http.HandleFunc("/server/post", middleware.IsAuth(middleware.CheckMethod(controller.PostHandler, "post")))
}
