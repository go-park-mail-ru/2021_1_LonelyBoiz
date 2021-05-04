package delivery

import (
	"google.golang.org/grpc/status"
	"net/http"
	model "server/internal/pkg/models"
	userProto "server/internal/user_server/delivery/proto"
)

func (a *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	//cookieId, ok := a.UserCase.GetIdFromContext(r.Context())
	//if !ok {
	//	response := model.ErrorResponse{Err: model.SessionErrorDenAccess}
	//	model.Process(model.LoggerFunc(response.Err, a.UserCase.LogInfo), model.ResponseFunc(w, 403, response))
	//	return
	//}
	//
	//vars := mux.Vars(r)
	//userId, err := strconv.Atoi(vars["id"])
	//if err != nil {
	//	response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
	//	response.Description["id"] = "Юзера с таким id нет"
	//	model.Process(model.LoggerFunc(response.Err, a.UserCase.LogInfo), model.ResponseFunc(w, 400, response))
	//	return
	//}
	//
	//if cookieId != userId {
	//	response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Отказано в доступе"}
	//	response.Description["id"] = "Пытаешься удалить не себя"
	//	model.Process(model.LoggerFunc(response.Err, a.UserCase.LogInfo), model.ResponseFunc(w, 403, response))
	//	return
	//}
	//
	//err = a.UserCase.DeleteUser(userId)
	//if err != nil {
	//	model.ResponseWithJson(w, 500, nil)
	//	model.Process(model.LoggerFunc(err.Error(), a.UserCase.LogError), model.ResponseFunc(w, 500, nil))
	//	return
	//}
	//
	//_, err = a.Sessions.Delete(r.Context(), &sessionproto.SessionId{Id: int32(cookieId)})
	//if err != nil {
	//	model.ResponseWithJson(w, 500, nil)
	//	model.Process(model.LoggerFunc(err.Error(), a.UserCase.LogError), model.ResponseFunc(w, 500, nil))
	//	return
	//}
	_, err := a.Server.DeleteUser(r.Context(), &userProto.UserNothing{})
	if err != nil {
		st, _ := status.FromError(err)
		if st.Code() != 200 {
			model.Process(model.LoggerFunc(st.Message(), a.UserCase.LogError), model.ResponseFunc(w, int(st.Code()), st.Message()))
			return
		}
	}

	cookie, _ := r.Cookie("token")

	a.UserCase.DeleteSession(cookie)
	http.SetCookie(w, cookie)

	a.UserCase.LogInfo(cookie.Expires)
	model.ResponseFunc(w, 200, nil)
}
