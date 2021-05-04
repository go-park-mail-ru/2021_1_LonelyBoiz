package delivery

import (
	"google.golang.org/grpc/status"
	"net/http"
	model "server/internal/pkg/models"
)

func (a *UserHandler) ChangeUserInfo(w http.ResponseWriter, r *http.Request) {
	newUser, err := a.UserCase.ParseJsonToUser(r.Body)
	if err != nil {
		response := model.ErrorResponse{Err: "Не удалось прочитать тело запроса"}
		model.Process(model.LoggerFunc(response.Err, a.UserCase.LogError), model.ResponseFunc(w, 400, response))
		return
	}

	protoUser, ok := a.UserCase.User2ProtoUser(newUser)
	if !ok {
		model.Process(model.LoggerFunc("Proto User Error", a.UserCase.LogError), model.ResponseFunc(w, 500, nil))
	}

	if protoUser == nil {
		response := model.ErrorResponse{Err: "Не удалось прочитать тело запроса"}
		model.Process(model.LoggerFunc(response.Err, a.UserCase.LogError), model.ResponseFunc(w, 400, response))
		return
	}

	user, err := a.Server.ChangeUser(r.Context(), protoUser)
	if err != nil {
		st, _ := status.FromError(err)
		model.Process(model.LoggerFunc(st.Message(), a.UserCase.LogError), model.ResponseFunc(w, int(st.Code()), st.Message()))
		return
	}

	nUser, ok := a.UserCase.ProtoUser2User(user)
	if !ok {
		model.Process(model.LoggerFunc("Proto User Error", a.UserCase.LogError), model.ResponseFunc(w, 500, nil))
	}

	model.Process(model.LoggerFunc("Success Change User Info", a.UserCase.LogInfo), model.ResponseFunc(w, 200, nUser))
}
