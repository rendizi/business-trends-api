package handler

import (
	"bta/internal/chatgpt"
	"bta/internal/db"
	"bta/internal/server"
	"net/http"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
	login := server.GetLogin(w, r)
	if login == "" {
		return
	}

	user, err := db.GetUser(login)
	if err != nil {
		server.Error(map[string]interface{}{"message": err.Error(), "status": 400}, w)
		return
	}
	url, err := chatgpt.Send(user)
	if err != nil {
		server.Error(map[string]interface{}{"message": err.Error(), "status": 400}, w)
		return
	}
	server.Ok(map[string]interface{}{"url": url}, w)
}
