package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

func (a *App) validateCookie(cookie string) bool {
	var mutex = &sync.Mutex{}
	mutex.Lock()
	defer mutex.Unlock()
	for _, userSessions := range a.Sessions {
		for _, v := range userSessions {
			if v.Value == cookie {

				return true
			}
		}
	}

	return false
}

func (a *App) GetUserInfo(w http.ResponseWriter, r *http.Request) {
	token, err := r.Cookie("token")
	if err != nil {
		w.WriteHeader(401)
		response := errorResponse{map[string]string{}, "Не залогинен"}
		json.NewEncoder(w).Encode(response)
		return
	}

	if !a.validateCookie(token.Value) {
		w.WriteHeader(401)
		response := errorResponse{map[string]string{}, "Кука устарела"}
		json.NewEncoder(w).Encode(response)
		return
	}

	userId, err := strconv.Atoi(strings.SplitAfter(r.URL.String(), "/")[2])
	if err != nil {
		w.WriteHeader(500)
	}

	var mutex = &sync.Mutex{}
	mutex.Lock()
	response := a.Users[userId]
	mutex.Unlock()

	response.PasswordHash = nil
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(response)

	fmt.Println("successful get user")
}

/*
curl -b 'token=abcdef' http://localhost:8001/users/1
*/
