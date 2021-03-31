package api

import (
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/sha3"
	"log"
	"net/http"
	"sync"
	"time"
)

type errorDescriptionResponse struct {
	Description map[string]string `json:"description"`
	Err         string            `json:"error"`
}

type errorResponse struct {
	Err string `json:"error"`
}

func (e errorDescriptionResponse) Error() string {
	ret, _ := json.Marshal(e)

	return string(ret)
}

func responseWithJson(w http.ResponseWriter, code int, body interface{}) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(body)
}

func validateSignUpData(newUser User) error {
	response := errorDescriptionResponse{Description: map[string]string{}, Err: "Не удалось зарегистрироваться"}

	_, err := govalidator.ValidateStruct(newUser)
	if err != nil {
		response.Description = govalidator.ErrorsByField(err)

		if newUser.Password != newUser.SecondPassword {
			response.Description["password"] = "Пароли не совпадают"
		}

		return response
	}

	if newUser.Password != newUser.SecondPassword {
		response.Description["password"] = "Пароли не совпадают"
		return response
	}

	return nil
}

func (a *App) isAlreadySignedUp(newEmail string) (bool, error) {
	var mutex = &sync.Mutex{}
	mutex.Lock()
	defer mutex.Unlock()
	for _, v := range a.Users {
		if v.Email == newEmail {
			response := errorDescriptionResponse{Description: map[string]string{}, Err: "Не удалось зарегистрироваться"}
			response.Description["mail"] = "Почта уже зарегистрирована"
			return true, response
		}
	}

	return false, nil
}

func hashPassword(pass string) ([]byte, error) {
	firstHash := sha3.New512()
	firstHash.Write([]byte(pass))
	secondHash, err := bcrypt.GenerateFromPassword(firstHash.Sum(nil), 14)
	return secondHash, err
}

func (a *App) addNewUser(newUser *User) error {
	var err error
	newUser.PasswordHash, err = hashPassword(newUser.Password)
	if err != nil {
		return err
	}
	newUser.Password = ""
	newUser.SecondPassword = ""

	var mutex = &sync.Mutex{}
	mutex.Lock()
	defer mutex.Unlock()
	newUser.Id = a.UserIds
	a.UserIds++
	a.Users[newUser.Id] = *newUser

	return nil
}

func (a *App) setSession(w http.ResponseWriter, id int) {
	key := KeyGen()
	expiration := time.Now().Add(24 * time.Hour)

	cookie := http.Cookie{
		Name:     "token",
		Value:    key,
		Expires:  expiration,
		SameSite: http.SameSiteNoneMode,
		//Secure:   true,
		Domain:   "localhost:3000",
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)

	var mutex = &sync.Mutex{}
	mutex.Lock()
	a.Sessions[id] = append(a.Sessions[id], cookie)
	mutex.Unlock()
}

func (a *App) SignUp(w http.ResponseWriter, r *http.Request) {
	var newUser User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	defer r.Body.Close()
	if err != nil {
		response := errorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
		responseWithJson(w, 400, response)
		return
	}

	if response := validateSignUpData(newUser); response != nil {
		responseWithJson(w, 400, response)
		return
	}

	if isSignedUp, response := a.isAlreadySignedUp(newUser.Email); isSignedUp {
		responseWithJson(w, 400, response)
		return
	}

	err = a.addNewUser(&newUser)
	if err != nil {
		responseWithJson(w, 500, nil)
		return
	}

	a.setSession(w, newUser.Id)

	newUser.Password = ""
	newUser.SecondPassword = ""
	newUser.PasswordHash = nil
	responseWithJson(w, 200, newUser)

	log.Println("successful registration\n", newUser)
}
