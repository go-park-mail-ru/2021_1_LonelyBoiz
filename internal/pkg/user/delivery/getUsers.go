package delivery

//
//import (
//	"encoding/json"
//	"log"
//	"net/http"
//	"server/api"
//	model "server/internal/pkg/models"
//)
//
////func (a *App) listUsers(newUser model.User) []model.User {
////	//var mutex = &sync.Mutex{}
////	//mutex.Lock()
////	//users := a.Users
////	//mutex.Unlock()
////
////	var usersRet []model.User
////
////	for _, v := range users {
////		if v.Id == newUser.Id {
////			continue
////		}
////
////		if (v.DatePreference == "both" || v.DatePreference == newUser.Sex) &&
////			(newUser.DatePreference == "both" || newUser.DatePreference == v.Sex) {
////			v.PasswordHash = nil
////			usersRet = append(usersRet, v)
////		}
////
////		if len(usersRet) == 5 {
////			break
////		}
////	}
////
////	return usersRet
////}
//
//func (a *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
//	var newUser model.User
//	decoder := json.NewDecoder(r.Body)
//	err := decoder.Decode(&newUser)
//	defer r.Body.Close()
//
//	if err != nil {
//		response := errorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
//		responseWithJson(w, 400, response)
//		return
//	}
//
//	ctx := r.Context()
//	id, ok := ctx.Value(ctxUserId).(int)
//	if !ok {
//		log.Println("error: get id from context")
//	}
//
//	if id != newUser.Id {
//		response := errorResponse{Err: "Отказано в доступе, кука устарела"}
//		responseWithJson(w, 401, response)
//		return
//	}
//
//	if !ValidateSex(newUser.Sex) {
//		response := errorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
//		response.Description["sex"] = "Неверно введен пол"
//		responseWithJson(w, 400, response)
//		return
//	}
//
//	if !ValidateDatePreferensces(newUser.DatePreference) {
//		response := errorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
//		response.Description["datePreferences"] = "Неверно введены предпочтения"
//		responseWithJson(w, 400, response)
//		return
//	}
//
//	//response := a.listUsers(newUser)
//
//	//responseWithJson(w, 200, response)
//	//
//	//log.Println("successful get user", response)
//}
