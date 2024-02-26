package controller



// func WsGroupHandler(w http.ResponseWriter, r *http.Request) {
// 	_, _, UserID := helper.Auth(DB, r)
// 	user := models.User{ID: UserID}
// 	user.GetUserById(DB, UserID)

// 	groupId := r.URL.Query().Get("groupId")
// 	grip, err := strconv.Atoi(groupId)
// 	fmt.Print(groupId)

// 	fmt.Println("received a request YYYYYYYYYYYYYEEEEEEEEEESSSSSSSSSSSS")
// 	conn, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	// defer conn.Close()
// 	client := &Client{Conn: conn, FirstName: user.FirstName, LastName: user.LastName, ID: user.ID, group_id: grip}
// 	Clients[client] = true
// 	fmt.Println("NEW CLIENT ", client)
// 	// fmt.Println("the new client is ", client.Session.Username)
// 	// mes, _ := json.Marshal(ConnexionMsg{Action: "connexion", Username: user.FirstName+" "+user.LastName})
// 	// PreventAllUsers(mes)
// 	go client.ListenGroup()

// }
