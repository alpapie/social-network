package controller

import "net/http"

func Home(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(500)
	w.Write([]byte(`{"error":"Wrong methode","success":false}`))
}