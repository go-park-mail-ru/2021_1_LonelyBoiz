package delivery

import (
	"log"
	"net/http"
	model "server/internal/pkg/models"
)

//func (a *App) listUsers(newUser model.User) []model.User {
//	//var mutex = &sync.Mutex{}
//	//mutex.Lock()
//	//users := a.Users
//	//mutex.Unlock()
//
//	var usersRet []model.User
//
//	for _, v := range users {
//		if v.Id == newUser.Id {
//			continue
//		}
//
//		if (v.DatePreference == "both" || v.DatePreference == newUser.Sex) &&
//			(newUser.DatePreference == "both" || newUser.DatePreference == v.Sex) {
//			v.PasswordHash = nil
//			usersRet = append(usersRet, v)
//		}
//
//		if len(usersRet) == 5 {
//			break
//		}
//	}
//
//	return usersRet
//}

func (a *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	newUser, err := a.UserCase.ParseJsonToUser(r.Body)
	if err != nil {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
		model.ResponseWithJson(w, 400, response)
		return
	}

	ctx := r.Context()
	id, ok := ctx.Value(model.CtxUserId).(int)
	if !ok {
		log.Println("error: get id from context")
	}

	if id != newUser.Id {
		response := model.ErrorResponse{Err: "Отказано в доступе, кука устарела"}
		model.ResponseWithJson(w, 401, response)
		return
	}

	if !a.UserCase.ValidateSex(newUser.Sex) {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
		response.Description["sex"] = "Неверно введен пол"
		model.ResponseWithJson(w, 400, response)
		return
	}

	if !a.UserCase.ValidateDatePreferences(newUser.DatePreference) {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
		response.Description["datePreferences"] = "Неверно введены предпочтения"
		model.ResponseWithJson(w, 400, response)
		return
	}

	//response := a.listUsers(newUser)
	//
	//responseWithJson(w, 200, response)
	//
	//
	//log.Println("successful get user", response)
	return
}
