package helper

import (
	"database/sql"
	"encoding/json"
	"fmt"

	// "forum/models"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func SessionAddOrUpdate(db *sql.DB, sssid, useremail string, user_id int) error {
	req := `SELECT uiSession,email,datefin from Session Where email='` + useremail + `';`
	var sessionid, email string
	var datef time.Time
	var errsession error
	_ = db.QueryRow(req).Scan(&sessionid, &email, &datef)

	if email == useremail {
		_, errsession = db.Exec("UPDATE Session SET uiSession=?, datefin=? where email=?;", sssid, time.Now().Add(time.Hour*24*3), email)
	} else {
		_, errsession = db.Exec("INSERT INTO Session (User_id,uiSession,email,datefin) VALUES(?,?,?,?);", user_id, sssid, useremail, time.Now().Add(time.Hour*24*3))
	}
	return errsession
}

func Auth(Db *sql.DB, r *http.Request) (bool, string, int) {

	sessionpi, err := r.Cookie("sessionId")
	if err != nil || sessionpi.String() == "" {
		return false, "", 0
	}
	fmt.Println(sessionpi.Value)

	var sessionId, email string
	var datef time.Time
	var user_id int
	req := `SELECT User_id, uiSession,email,datefin from Session Where uiSession=?`
	err = Db.QueryRow(req, sessionpi.Value).Scan(&user_id, &sessionId, &email, &datef)

	if err != nil {
		return false, "", 0
	}

	fmt.Println(email)

	if sessionId != "" && email != "" && datef.After(time.Now()) {
		return true, email, user_id
	}
	return false, "", 0
}

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
	return strings.EqualFold(methode, r.Method)
}

func DeleteSessio(db *sql.DB, ssid string) error {
	req := `DELETE from Session Where uiSession=?;`
	_, err := db.Exec(req, ssid)
	return err
}

// ******************************* PARSE FILE IN URL *****************
func PArseUlr(r *http.Request, match string) (bool, int) {
	index := strings.Split(r.URL.Path[1:], "/")
	fmt.Println(index[0] + "/" + index[1])
	fmt.Println(len(index))
	fmt.Println(match)
	if len(index) == 3 && index[0]+"/"+index[1] == match {
		id, err := strconv.Atoi(index[2])
		if err == nil {
			return true, id
		}
	}
	return false, 0
}

func FecthError(ch []error) bool {
	for _, err := range ch {
		if err != nil {
			fmt.Println(err)
			return true
		}
	}
	return false
}

func ParseCatId(cat []string) ([]int, error) {
	catid := []int{}
	for _, v := range cat {
		a, errt := strconv.Atoi(v)
		if errt != nil {
			return []int{}, errt
		}
		catid = append(catid, a)
	}
	return catid, nil
}

func WriteJSON(w http.ResponseWriter, status int, data map[string]interface{}, headers http.Header) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}
	js = append(js, '\n')
	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Write(js)
	return nil
}
