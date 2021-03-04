package api

import (
	"fmt"
	"io"
	"net/http"
)

func (a *App) LogOut(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	io.WriteString(w, "Log out")
}
