package middleware

import (
	// "social_network/controllers"
	"fmt"
	"net/http"
	"social_network/controller"
	helper "social_network/helper"
	"strings"
)

func CheckRequest(r *http.Request, path, methode string) (bool, int) {
	if strings.ToLower(r.Method) == methode && r.URL.Path == path {
		return true, 0
	} else if !Getmethode(r, methode) {
		return false, 405
	} else {
		return false, 404
	}
}

func Getmethode(r *http.Request, methode string) bool {
	if strings.ToLower(r.Method) != methode {
		return false
	}
	return true
}

// *****************************CHECK IF THE USER IS LOGGING OR NOT***************************
func Log(next http.HandlerFunc) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		
		if is, _,_ := helper.Auth(controller.DB, r); is {
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte(`{"error":"no access","success":false}`))
		}
	}
	return http.HandlerFunc(fn)
}

func IsAuth(next http.HandlerFunc) http.HandlerFunc {
	fnt := func(w http.ResponseWriter, r *http.Request) {
		if is, _ ,_:= helper.Auth(controller.DB, r); is {
			w.WriteHeader(302)
			w.Write([]byte(`{"error":"Connected","success":false}`))
		} else {
			next.ServeHTTP(w, r)
		}
	}
	return http.HandlerFunc(fnt)
}

func Ispath(next http.HandlerFunc, path string) http.HandlerFunc {
	fnt := func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/server/"+path {
			next.ServeHTTP(w, r)
		} else {
			fmt.Println("error 404")
			w.WriteHeader(404)
			w.Write([]byte(`{"error":"page not found","success":false}`))
		}
	}
	return http.HandlerFunc(fnt)
}

func CheckMethod(next http.HandlerFunc, methode string) http.HandlerFunc {
	fnt := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method)
		if strings.ToLower(r.Method) == methode {
			next.ServeHTTP(w, r)
			return
		}
		w.WriteHeader(405)
		w.Write([]byte(`{"error":"Wrong methode","success":false}`))
	}
	return http.HandlerFunc(fnt)
}
