package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/sha3"
)

func (a *App) validateCookieForChanging(cookie string, id int) bool {
	for _, v := range a.Sessions[id] {
		fmt.Println(v.Value, cookie)
		if v.Value == cookie {
			return true
		}
	}

	return false
}

func (a *App) checkPasswordForCHanging(newUser *User) bool {
	for _, v := range a.Users {
		if v.Email == newUser.Email {
			pass := sha3.New512()
			pass.Write([]byte(newUser.OldPass))
			err := bcrypt.CompareHashAndPassword(v.PasswordHash, pass.Sum(nil))
			if err != nil {
				return false
			}

			newUser.Id = v.Id
			return true
		}
	}

	return false
}

/*func (a *App) validatePass() bool {

}*/

func (a *App) ChangeUserInfo(w http.ResponseWriter, r *http.Request) {
	var newUser User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		w.WriteHeader(400)
		response := errorResponse{map[string]string{}, "Неверный запрос"}
		json.NewEncoder(w).Encode(response)
		return
	}

	token, err := r.Cookie("token")
	if err != nil {
		w.WriteHeader(401)
		response := errorResponse{map[string]string{}, "Не залогинен 1"}
		json.NewEncoder(w).Encode(response)
		return
	}

	userId, err := strconv.Atoi(strings.SplitAfter(r.URL.String(), "/")[2])
	if err != nil {
		w.WriteHeader(500)
	}
	if !a.validateCookieForChanging(token.Value, userId) {
		w.WriteHeader(401)
		response := errorResponse{map[string]string{}, "Кука устарела"}
		json.NewEncoder(w).Encode(response)
		return
	}

	fmt.Println(newUser)

	/*var newUser User
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&newUser)
	if err != nil {
		w.WriteHeader(400)
		response := errorResponse{map[string]string{}, "Неверный запрос"}
		json.NewEncoder(w).Encode(response)
		return
	}

	response := a.Users[userId]
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(response)

	fmt.Println("successful get user")
	*/
}

/*
curl --header "Content-Type: application/json" \
  	 --request PATCH \
  	 --data '{"mail":"xyz","pass":"xyz","passRepeat":"xyz","oldPass":"xyz1","name":"vasya"}' \
  	 http://localhost:8002/users
*/
