package models

import (
	"database/sql"
)

func (u *User) GetFollow(DB *sql.DB, limit int) ([]User, error) {
	req := `select id,firstName,lastName,email,nickName from User `
	row, err := DB.Query(req)
	users := []User{}
	if err != nil {
		return users, err
	}

	for row.Next() {
		newUser := User{}
		row.Scan(&newUser.ID, &newUser.FirstName, &newUser.LastName, &newUser.Email, &newUser.NickName)
		users = append(users, newUser)
	}
	defer row.Close()
	return users, nil
}

