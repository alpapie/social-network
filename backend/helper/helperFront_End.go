package helper

import (
	"html/template"
	"net/http"
	"strconv"
)

func RenderTemplate(w http.ResponseWriter, s string, Data interface{}) error {
	page, err := template.ParseFiles(s + ".html")
	if err != nil {
		return err
	}
	return page.Execute(w, Data)
}

func ErrorPage(w http.ResponseWriter, i int) error {
	DataError := struct {
		Code    string
		Message string
	}{
		Code:    strconv.Itoa(i),
		Message: http.StatusText(i),
	}

	w.WriteHeader(i)
	return WriteJSON(w, i, map[string]interface{}{"error": DataError,"success":false}, nil)
}

func ErrorMessage(w http.ResponseWriter,message string,) error{
	w.WriteHeader(200)
	return WriteJSON(w, 0, map[string]interface{}{"error": message,"success":false,}, nil)
}

// ******************* VERIF IF THE STRING IS AN INT*****************************************************
func IsInt(s string) bool {
	for _, v := range s {
		if v < '0' || v > '9' {
			return false
		}
	}
	return true
}
