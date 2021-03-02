package api

import (
	"fmt"
	"io"
	"net/http"
)

func (a *App) SignUp(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	io.WriteString(w, "Sign Un")
}
