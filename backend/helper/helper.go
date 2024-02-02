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
func SessionAddOrUpdate(db *sql.DB, sssid, useremail string) error {
	req := `SELECT sessionId,email,datefin from Session Where email='` + useremail + `';`
	// req:=fmt.Sprintf(`SELECT * from Session Where email=?;`)
	row, err := db.Query(req)
	var sessionid, email string
	var datef time.Time
	var errsession error
	if err != nil {
		fmt.Println(err)
		return err
	}

	for row.Next() {
		row.Scan(&sessionid, &email, &datef)

	}
	
	if email == useremail {
		_, errsession = db.Exec("UPDATE Session SET sessionId=?, datefin=? where email=?;", sssid, time.Now().Add(time.Hour*24*3) ,email)
	} else {
		_, errsession = db.Exec("INSERT INTO Session (sessionId,email,datefin) VALUES(?,?,?);", sssid, useremail, time.Now().Add(time.Hour*24*3))
	}
	return errsession

}

func Auth(Db *sql.DB, r *http.Request) (bool, string) {

	sessionpi, err := r.Cookie("sessionid")
	if err != nil || sessionpi.String() == "" {
		return false, ""
	}
	var Id int
	var sessionId, email string
	var datef time.Time
	req := `SELECT * from Session Where sessionId=?;`
	row, err := Db.Query(req, sessionpi.Value)

	if err != nil {
		return false, ""
	}
	for row.Next() {
		row.Scan(&Id, &sessionId, &email, &datef)
	}

	if sessionId != "" && email != "" && datef.After(time.Now()) {
		return true, email
	}
	return false, ""
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
	if strings.ToLower(r.Method) != methode {
		return false
	}
	return true
}
func DeleteSessio(db *sql.DB, ssid string) error {
	req := `DELETE from Session Where sessionId=?;`
	_, err := db.Exec(req, ssid)
	return err
}

// ******************************* PARSE FILE IN URL *****************
func PArseUlr(r *http.Request, match string) (bool, int) {
	index := strings.Split(r.URL.Path[1:], "/")
	fmt.Println(index[0]+"/"+index[1])
	fmt.Println(len(index))
	fmt.Println(match)
	if len(index) == 3 && index[0]+"/"+index[1]== match {
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

func ErrorMessage(w http.ResponseWriter,message string,) error{
	w.WriteHeader(200)
	return WriteJSON(w, 0, map[string]interface{}{"error": message,"success":false,}, nil)
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
	w.Write(js)

	return nil
}

