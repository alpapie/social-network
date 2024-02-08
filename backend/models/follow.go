package models

import "database/sql"

func (u *User)GetFollow(DB *sql.DB, limit int)([]User,error){
	req:=`select id,firstName,lastName,email,nickName fom User `
	row,err:=DB.Query(req)
	users:=[]User{}
	if err!=nil {
		return users,err
	}
	for row.Next() {
		row.Scan(&u.ID,&u.FirstName,&u.LastName,&u.Email,&u.NickName)
		users = append(users, *u)
	}
	defer row.Close()
	return users,nil
}