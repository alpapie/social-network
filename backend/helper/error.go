package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"social_network/models"
)

func Error(w http.ResponseWriter, er error, statusCode int) {
	fmt.Println("Error ", statusCode, er)
	http.Error(w, "Error ", statusCode)
}

func WriteError(w http.ResponseWriter, messageErr string, statusCode int) {
	fmt.Println(messageErr)

	w.Header().Set("Content-Type", "application/json")
	switch statusCode {
	case 500:
		w.WriteHeader(http.StatusInternalServerError)
		err := parseError(w, "Internal Server Error", statusCode)
		if err != nil {
			fmt.Println("my error", err)
		}
	case 400:
		w.WriteHeader(http.StatusBadRequest)
		err := parseError(w, messageErr, statusCode)
		if err != nil {
			fmt.Println("my error", err)
		}

	case 404:
		w.WriteHeader(http.StatusNotFound)
		err := parseError(w, "Not Found", statusCode)
		if err != nil {
			fmt.Println("my error", err)
		}

	case 405:
		fmt.Println("Method not allowed")
		w.WriteHeader(http.StatusMethodNotAllowed)

		err := parseError(w, "Method Not Allowed", statusCode)
		if err != nil {
			fmt.Println("my error", err)
		}

	case 401:
		w.WriteHeader(http.StatusUnauthorized)
		err := parseError(w, "Unauthorized", statusCode)
		if err != nil {
			fmt.Println("my error", err)
		}

		// default:
		// 	w.WriteHeader(http.StatusBadRequest)
	}
}

func parseError(w http.ResponseWriter, message string, code int) error {

	newError := models.Error{
		StatusCode: code,
		Message:    message,
	}

	jsonData, err := json.Marshal(newError)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return err
	}

	_, err = w.Write(jsonData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return err
	}
	return nil
}
