package api

import (
	"fmt"
	"io"
	"net/http"
)

func (a *App) DownloadPhoto(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	io.WriteString(w, "download photo")
}
