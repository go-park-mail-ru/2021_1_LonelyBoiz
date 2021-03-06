package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (a *App) SignUp(w http.ResponseWriter, r *http.Request) {
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
