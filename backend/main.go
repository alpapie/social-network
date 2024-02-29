package main

import (
	"fmt"
	"net/http"
	"os"
	conf "social_network/config"
	"social_network/controller"
	migr "social_network/db/sqlite"
	route "social_network/routes"
)

var port = ":8080"

func init() {
	var err error
	controller.DB, err = conf.GetDB()
	if err != nil {
		fmt.Println("connection database Error", err)
		os.Exit(0)
	}
	if err_migr := migr.ApplyMigrations(); err_migr != nil {
		//rendre un json pour dire que le server est down
		os.Exit(0)
	}
}
func main() {
	fmt.Println("Listening in http://localhost" + port)
	route.Routes()
	http.ListenAndServe(port, nil)
}
