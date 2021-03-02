package api

import (
	"fmt"
	"io"
	"net/http"
)

func (a *App) SignIn(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	io.WriteString(w, "Sign In")
}
