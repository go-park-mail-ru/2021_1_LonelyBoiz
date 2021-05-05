package delivery

import (
	"net/http"
)

func (a *UserHandler) GetSecreteAlbum(w http.ResponseWriter, r *http.Request) {
	//getterId, ok := a.Sessions.GetIdFromContext(r.Context())
	//if !ok {
	//	response := model.ErrorResponse{Err: model.SessionErrorDenAccess}
	//	model.Process(model.LoggerFunc(response.Err, a.UserCase.LogInfo), model.ResponseFunc(w, 403, response))
	//	return
	//}
	//
	//vars := mux.Vars(r)
	//ownerId, err := strconv.Atoi(vars["ownerId"])
	//if err != nil {
	//	response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Неправильные входные данные"}
	//	response.Description["id"] = "Пользователя с таким id нет"
	//	model.Process(model.LoggerFunc(response.Err, a.UserCase.LogInfo), model.ResponseFunc(w, 400, response))
	//	return
	//}
	//
	//var photos []string
	//photos, code, err := a.UserCase.GetSecreteAlbum(ownerId, getterId)
	//if code == 500 {
	//	model.Process(model.LoggerFunc(err, a.UserCase.LogError), model.ResponseFunc(w, code, nil))
	//	return
	//}
	//
	//res := make(map[string][]string, 1)
	//
	//res["photos"] = photos
	//
	//w.WriteHeader(code)
	//json.NewEncoder(w).Encode(res)
	//
	//a.UserCase.LogInfo("Success get secrete album")
}
