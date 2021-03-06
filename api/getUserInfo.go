package api

import (
	"fmt"
	"io"
	"net/http"
)

func (a *App) GetUserInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	w.WriteHeader(200)
	io.WriteString(w, "\nget users info")
}
