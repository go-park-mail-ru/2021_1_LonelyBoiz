package delivery

import "net/http"

type PhotoDeliveryInterface interface {
	UploadPhoto(w http.ResponseWriter, r *http.Request)
	DownloadPhoto(w http.ResponseWriter, r *http.Request)
}
