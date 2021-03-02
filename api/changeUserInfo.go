package api

import (
	"fmt"
	"io"
	"net/http"
)

func (a *App) ChangeUserInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	io.WriteString(w, "Change user info")
}
