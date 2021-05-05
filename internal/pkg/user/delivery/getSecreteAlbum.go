package delivery

import (
	"encoding/json"
	"net/http"
	model "server/internal/pkg/models"
	userProto "server/internal/user_server/delivery/proto"

	"google.golang.org/grpc/status"
)

func (a *UserHandler) GetSecreteAlbum(w http.ResponseWriter, r *http.Request) {
	a.UserCase.LogInfo("Передано на сервер USER")
	photos, err := a.Server.GetSecreteAlbum(r.Context(), &userProto.UserNothing{})
	if err != nil {
		st, _ := status.FromError(err)
		model.Process(model.LoggerFunc(st.Message(), a.UserCase.LogError), model.ResponseFunc(w, int(st.Code()), st.Message()))
		return
	}
	a.UserCase.LogInfo("Получен результат из сервера USER")

	res := make(map[string][]string, 1)

	res["photos"] = a.UserCase.ProtoPhotos2Photos(photos.Photos)
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(res)

	a.UserCase.LogInfo("Success get secrete album")
}
