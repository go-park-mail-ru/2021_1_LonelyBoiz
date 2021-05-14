package delivery

import (
	"encoding/json"
	"net/http"
	"server/internal/pkg/models"
	userproto "server/internal/user_server/delivery/proto"

	"google.golang.org/grpc/status"
)

func (a *UserHandler) LikesHandler(w http.ResponseWriter, r *http.Request) {
	var like models.Like
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&like)
	defer r.Body.Close()
	if err != nil {
		response := models.ErrorResponse{Err: "Не удалось прочитать тело запроса"}
		models.Process(models.LoggerFunc(response.Err, a.UserCase.LogError), models.ResponseFunc(w, 400, response))
		return
	}

	a.UserCase.LogInfo("Передано на сервер USER")
	chat, err := a.Server.CreateChat(r.Context(), &userproto.Like{
		UserId:   int32(like.UserId),
		Reaction: like.Reaction,
	})
	a.UserCase.LogInfo("Получен результат из сервера USER")

	if err != nil {
		st, _ := status.FromError(err)
		if st.Code() == 204 {
			w.WriteHeader(204)
			return
		}
		models.Process(models.LoggerFunc(st.Message(), a.UserCase.LogError), models.ResponseFunc(w, int(st.Code()), st.Message()))
		return
	}

	nChat := models.Chat{
		ChatId:              int(chat.GetChatId()),
		PartnerId:           int(chat.GetPartnerId()),
		PartnerName:         chat.GetPartnerName(),
		LastMessage:         chat.GetLastMessage(),
		LastMessageTime:     chat.GetLastMessageTime(),
		LastMessageAuthorId: int(chat.GetLastMessageAuthorId()),
		Photos:              a.UserCase.ProtoPhotos2Photos(chat.GetPhotos()),
	}

	models.Process(models.LoggerFunc("Create Feed", a.UserCase.LogInfo), models.ResponseFunc(w, 200, nChat))

	a.UserCase.WebsocketChat(&nChat)
}
