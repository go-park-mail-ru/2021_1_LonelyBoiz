package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

type key int

const ctxUserId key = -1

func (a *App) ValidateCookie(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := r.Cookie("token")
		if err != nil {
			response := errorResponse{Err: "Вы не авторизованы"}
			responseWithJson(w, 401, response)
			return
		}

		//здесь будет поход в базу
		var mutex = &sync.Mutex{}
		mutex.Lock()
		defer mutex.Unlock()
		for id, userSessions := range a.Sessions {
			for _, v := range userSessions {
				if v.Value == token.Value {
					ctx := r.Context()
					ctx = context.WithValue(ctx,
						ctxUserId,
						10,
					)
					fmt.Println("id=", id)
					fmt.Println("валидационная мидалварь")
					next.ServeHTTP(w, r.WithContext(ctx))
					return
				}
			}
		}

		response := errorResponse{Err: "Вы не авторизованы"}
		responseWithJson(w, 401, response)
		return
	})
}

func (a *App) GetUserInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	if err != nil {
		response := errorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
		response.Description["id"] = "Пользоватея с таким id нет"
		responseWithJson(w, 400, response)
		return
	}

	ctx := r.Context()
	val, ok := ctx.Value(ctxUserId).(int)
	if !ok {
		fmt.Println("ne ok")
	}
	fmt.Println("val=", val)

	if val != userId {
		fmt.Println("sdfg")
		response := errorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
		response.Description["id"] = "Пользоватея с таким id нет"
		responseWithJson(w, 400, response)
		return
	}

	var mutex = &sync.Mutex{}
	mutex.Lock()
	userInfo, ok := a.Users[userId]
	mutex.Unlock()
	if !ok {
		response := errorDescriptionResponse{Description: map[string]string{}, Err: "Отказано в доступе, кука устарела"}
		response.Description["id"] = "Пользователя с таким id нет"
		responseWithJson(w, 400, response)
		return
	}

	userInfo.PasswordHash = nil
	responseWithJson(w, 200, userInfo)

	log.Println("successful get user", userInfo)
}
