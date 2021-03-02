package api

import (
	"fmt"
	"io"
	"net/http"
)

func (a *App) GetUserInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	io.WriteString(w, "get users info")
}
