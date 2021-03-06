package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (a *App) isAlreadySignedIn(newEmail string) bool {
	for _, v := range a.Users {
		if v.Email == newEmail {
			return true
		}
	}

	return false
}

func (a *App) addNewUser(newUser User) {
	newUser.Id = a.UserIds
	a.UserIds++
	a.Users = append(a.Users, newUser)

}

func (a *App) SignIn(w http.ResponseWriter, r *http.Request) {
	var newUser User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		fmt.Println("ne decoditsya")
		io.WriteString(w, "err")
		return
	}

	if a.isAlreadySignedIn(newUser.Email) {
		fmt.Println("uzhe zaregan")
		return
	}

	a.Users = append(a.Users, newUser)

	fmt.Println("\n\n", newUser)
	io.WriteString(w, "huy")
}
