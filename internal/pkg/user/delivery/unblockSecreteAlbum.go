package delivery

import (
	"net/http"
)

func (a *UserHandler) UnblockSecreteAlbum(w http.ResponseWriter, r *http.Request) {
	//ownerId, ok := a.Sessions.GetIdFromContext(r.Context())
	//if !ok {
	//	response := model.ErrorResponse{Err: model.SessionErrorDenAccess}
	//	model.Process(model.LoggerFunc(response.Err, a.UserCase.LogInfo), model.ResponseFunc(w, 403, response))
	//	return
	//}
	//
	//vars := mux.Vars(r)
	//getterId, err := strconv.Atoi(vars["getterId"])
	//if err != nil {
	//	response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Неправильные входные данные"}
	//	response.Description["id"] = "Пользователя с таким id нет"
	//	model.Process(model.LoggerFunc(response.Err, a.UserCase.LogInfo), model.ResponseFunc(w, 400, response))
	//	return
	//}
	//
	//code, err := a.UserCase.UnblockSecreteAlbum(ownerId, getterId)
	//if code == 500 {
	//	model.Process(model.LoggerFunc(err, a.UserCase.LogError), model.ResponseFunc(w, code, nil))
	//}
	//
	//model.Process(model.LoggerFunc("Successful unlock", a.UserCase.LogInfo), model.ResponseFunc(w, 204, nil))
}
