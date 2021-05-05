package delivery

import (
	"encoding/json"
	"google.golang.org/grpc/status"
	"net/http"
	"server/internal/pkg/models"
	userproto "server/internal/user_server/delivery/proto"
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

	chat, err := a.Server.CreateChat(r.Context(), &userproto.Like{
		UserId:   int32(like.UserId),
		Reaction: like.Reaction,
	})

	if err != nil {
		st, _ := status.FromError(err)
		if st.Code() == 204 {
			w.WriteHeader(204)
			return
		}

		models.Process(models.LoggerFunc(st.Message(), a.UserCase.LogError), models.ResponseFunc(w, int(st.Code()), st.Message()))
		return
	}

	photos, ok := a.UserCase.ProtoPhotos2Photos(chat.GetPhotos())
	if !ok {
		models.Process(models.LoggerFunc("Proto error", a.UserCase.LogError), models.ResponseFunc(w, 500, nil))
	}

	nChat := models.Chat{
		ChatId:              int(chat.GetChatId()),
		PartnerId:           int(chat.GetPartnerId()),
		PartnerName:         chat.GetPartnerName(),
		LastMessage:         chat.GetLastMessage(),
		LastMessageTime:     chat.GetLastMessageTime(),
		LastMessageAuthorId: int(chat.GetLastMessageAuthorId()),
		Photos:              photos,
	}

	models.Process(models.LoggerFunc("Create Feed", a.UserCase.LogInfo), models.ResponseFunc(w, 200, nChat))

	a.UserCase.WebsocketChat(&nChat)
}
