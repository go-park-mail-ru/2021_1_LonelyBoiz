package api

import (
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/sha3"
	"log"
	"net/http"
	model "server/models"
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

func validateSignUpData(newUser model.User) error {
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
	isSignUp := a.Db.CheckMail(newEmail)
	if isSignUp == true {
		response := errorDescriptionResponse{Description: map[string]string{}, Err: "Не удалось зарегистрироваться"}
		response.Description["mail"] = "Почта уже зарегистрирована"
		return true, response
	}

	return false, nil
}

func hashPassword(pass string) ([]byte, error) {
	firstHash := sha3.New512()
	firstHash.Write([]byte(pass))
	secondHash, err := bcrypt.GenerateFromPassword(firstHash.Sum(nil), 14)
	return secondHash, err
}

func (a *App) addNewUser(newUser *model.User) error {
	var err error
	newUser.PasswordHash, err = hashPassword(newUser.Password)
	if err != nil {
		return err
	}
	newUser.Password = ""
	newUser.SecondPassword = ""

	log.Println("45")

	id, err := a.Db.AddUser(*newUser)
	if err != nil {
		response := errorDescriptionResponse{Description: map[string]string{}, Err: err.Error()}
		return response
	}

	log.Println("46")
	newUser.Id = id

	return nil
}

func (a *App) setSession(w http.ResponseWriter, id int) error {
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

	err := a.Db.AddCookie(id, key)
	if err != nil {
		response := errorDescriptionResponse{Description: map[string]string{}, Err: err.Error()}
		return response
	}

	return nil
}

func (a *App) SignUp(w http.ResponseWriter, r *http.Request) {
	var newUser model.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	defer r.Body.Close()
	if err != nil {
		response := errorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
		responseWithJson(w, 400, response)
		return
	}

	log.Println("5")

	if response := validateSignUpData(newUser); response != nil {
		responseWithJson(w, 400, response)
		return
	}

	log.Println("4")

	if isSignedUp, response := a.isAlreadySignedUp(newUser.Email); isSignedUp {
		responseWithJson(w, 400, response)
		return
	}

	log.Println("3")
	err = a.addNewUser(&newUser)
	if err != nil {
		responseWithJson(w, 500, nil)
		return
	}
	log.Println("2")

	err = a.setSession(w, newUser.Id)
	if err != nil {
		response := errorDescriptionResponse{Description: map[string]string{}, Err: err.Error()}
		responseWithJson(w, 500, response)
		return
	}

	log.Println("1")

	newUser.Password = ""
	newUser.SecondPassword = ""
	newUser.PasswordHash = nil
	responseWithJson(w, 200, newUser)

	log.Println("successful registration\n", newUser)
}
