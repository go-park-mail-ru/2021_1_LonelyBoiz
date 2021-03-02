package api

import (
	"fmt"
	"io"
	"net/http"
)

func (a *App) DeletePhoto(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	io.WriteString(w, "delete photo")
}
