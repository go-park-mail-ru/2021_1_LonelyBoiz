package api

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/sha3"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

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

func validatePass(pass string) error {
	response := errorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
	switch {
	case len(pass) < 8 || len(pass) > 64:
		response.Description["password"] = "Пароль должен содержать 8 символов"
		return response
	case pass == strings.ToLower(pass) || pass == strings.ToUpper(pass):
		response.Description["password"] = "Пароль должен состоять из символов разного регистра"
		return response
	case !strings.ContainsAny(pass, "1234567890"):
		response.Description["password"] = "Пароль должен содержать цифру"
		return response
	}

	return nil
}

func validateEmail(email string) bool {
	if len(email) < 3 && len(email) > 254 {
		return false
	}

	return emailRegex.MatchString(email)
}

func validateSignUpData(newUser User) error {
	err := validatePass(newUser.Password)
	if err != nil {
		return err
	}

	response := errorDescriptionResponse{Description: map[string]string{}, Err: "Не удалось зарегестрироваться"}

	tm := time.Unix(newUser.Birthday, 0)

	diff := time.Now().Sub(tm)

	switch {
	case !validateEmail(newUser.Email):
		response.Description["mail"] = "Почта не прошла валидацию"
	case newUser.Name == "":
		response.Description["name"] = "Введите имя"
	case newUser.Password != newUser.SecondPassword:
		response.Description["password"] = "Пароли не совпадают"
	case diff/24/365 < 18:
		response.Description["Birthday"] = "Вам должно быть 18"
	}

	if len(response.Description) != 0 {
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
			response := errorDescriptionResponse{Description: map[string]string{}, Err: "Не удалось зарегестрироваться"}
			response.Description["mail"] = "Почта уже зарегестрирована"
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
	cookie := http.Cookie{Name: "token", Value: key, SameSite: 0, Expires: expiration, Domain: "https://lepick.herokuapp.com", HttpOnly: true}

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

	log.Println("successful resistration\n", newUser)
}
