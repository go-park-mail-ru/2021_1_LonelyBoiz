package api

import (
	"fmt"
	"io"
	"net/http"
)

func (a *App) UploadPhoto(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	io.WriteString(w, "upload photo")
}
