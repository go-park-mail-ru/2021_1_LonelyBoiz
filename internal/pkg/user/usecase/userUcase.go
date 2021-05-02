package usecase

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	model "server/internal/pkg/models"
	"server/internal/pkg/user/repository"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/microcosm-cc/bluemonday"

	"github.com/asaskevich/govalidator"
	"github.com/gorilla/websocket"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/sha3"
)

type UserUseCaseInterface interface {
	ValidateSex(sex string) bool
	CheckPasswordWithId(passToCheck string, id int) (bool, error)
	ValidateDatePreferences(pref string) bool
	CheckPasswordWithEmail(passToCheck string, email string) (bool, error)
	ChangeUserProperties(newUser *model.User) error
	ChangeUserPassword(newUser *model.User) error
	ValidatePassword(password string) bool
	ValidateSignInData(newUser model.User) (bool, error)
	ValidateSignUpData(newUser model.User) error
	IsAlreadySignedUp(newEmail string) (bool, error)
	HashPassword(pass string) ([]byte, error)
	IsActive(newUser *model.User) error
	AddNewUser(newUser *model.User) error
	ParseJsonToUser(body io.ReadCloser) (model.User, error)

	SignIn(user model.User) (newUser model.User, code int, err error)
	GetUserInfoById(id int) (user model.User, err error)
	DeleteUser(id int) error
	CheckCaptch(token string) (bool, error)
	CreateNewUser(newUser model.User) (user model.User, code int, responseError error)
	ChangeUserInfo(newUser model.User, id int) (user model.User, code int, err error)
	CreateFeed(id int, limitInt int) ([]int, int, error)
	CreateChat(id int, like model.Like) (model.Chat, int, error)
	WebsocketChat(newChat *model.Chat)
	SetChat(ws *websocket.Conn, id int)
	SetCookie(token string) http.Cookie

	model.LoggerInterface
	GetIdFromContext(ctx context.Context) (int, bool)
}

type UserUsecase struct {
	Clients *map[int]*websocket.Conn
	Db      repository.UserRepositoryInterface
	model.LoggerInterface
	Sanitizer *bluemonday.Policy
}

var googleCaptchaSecret string = os.Getenv("DATABASE_URL")

func (u *UserUsecase) SetChat(ws *websocket.Conn, id int) {
	(*u.Clients)[id] = ws
}

func (u *UserUsecase) WebsocketChat(newChat *model.Chat) {
	newChatToSend, err := u.Db.GetNewChatById(newChat.ChatId, newChat.PartnerId)
	if err != nil {
		u.LogError("Не удалось составить чат : " + err.Error())
		return
	}

	client, ok := (*u.Clients)[newChat.PartnerId]
	if !ok {
		u.LogInfo("Пользователь с id = " + strconv.Itoa(newChat.PartnerId) + " не в сети")
		return
	}

	response := model.WebsocketResponse{ResponseType: "chat", Object: newChatToSend}

	err = client.WriteJSON(response)
	if err != nil {
		u.LogError("Не удалось отправить сообщение")
		return
	}
}

func (u *UserUsecase) CreateChat(id int, like model.Like) (model.Chat, int, error) {
	if like.Reaction != "like" && like.Reaction != "skip" {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
		response.Description["like"] = "неправильный формат реакции, ожидается skip или like"
		u.LogInfo(response.Err)
		return model.Chat{}, 400, response
	}

	rowsAffected, err := u.Db.Rating(id, like.UserId, like.Reaction)
	if err != nil {
		u.LogError(err)
		return model.Chat{}, 500, nil
	}
	if rowsAffected != 1 {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Отказано в доступе"}
		response.Description["userID"] = "Пытаешься поставить лайк человеку не со своей ленты"
		u.LogInfo(response.Err)
		return model.Chat{}, 403, response
	}

	reciprocity, err := u.Db.CheckReciprocity(like.UserId, id)
	if err != nil {
		u.LogError(err)
		return model.Chat{}, 500, nil
	}

	if !reciprocity || like.Reaction == "skip" {
		u.LogInfo("Return 204 header")
		return model.Chat{}, 204, nil
	}

	var newChat model.Chat
	newChat.ChatId, err = u.Db.CreateChat(id, like.UserId)
	if err != nil {
		u.LogError(err)
		return model.Chat{}, 500, nil
	}

	newChat, err = u.Db.GetNewChatById(newChat.ChatId, id)
	if err != nil {
		u.LogError(err)
		return model.Chat{}, 500, nil
	}

	return newChat, 200, nil
}

func (u *UserUsecase) CreateFeed(id int, limitInt int) ([]int, int, error) {
	feed, err := u.Db.GetFeed(id, limitInt)

	if err != nil {
		u.LogError(err)
		return nil, 500, nil
	}

	if len(feed) < limitInt {
		err = u.Db.CreateFeed(id)

		if err != nil {
			u.LogError(err)
			return nil, 500, nil
		}

		feed, err = u.Db.GetFeed(id, limitInt)
		if err != nil {
			u.LogError(err)
			return nil, 500, nil
		}
	}
	if len(feed) == 0 {
		err := u.Db.ClearFeed(id)
		if err != nil {
			u.LogError(err)
			return nil, 500, nil
		}
		feed, err = u.Db.GetFeed(id, limitInt)
		if err != nil {
			u.LogError(err)
			return nil, 500, nil
		}
	}

	if len(feed) == 0 {
		feed = make([]int, 0)
	}

	return feed, 200, nil
}

func (u *UserUsecase) ChangeUserInfo(newUser model.User, id int) (user model.User, code int, err error) {
	newUser.Id = id
	if newUser.Password != "" {
		err := u.ChangeUserPassword(&newUser)

		if err != nil {
			return user, 400, err
		}

		newUser.Password = ""
		newUser.OldPassword = ""
		newUser.SecondPassword = ""
	}

	newUser.Id = id
	err = u.ChangeUserProperties(&newUser)
	if err != nil {
		if reflect.TypeOf(err) != reflect.TypeOf(model.ErrorDescriptionResponse{}) {
			return user, 500, nil
		}
		return user, 400, err
	}

	newUser.PasswordHash = nil
	return newUser, 200, nil
}

func (u *UserUsecase) CreateNewUser(newUser model.User) (user model.User, code int, responseError error) {
	ok, err := u.CheckCaptch(newUser.CaptchaToken)
	if err != nil {
		u.LogError(err)
		return model.User{}, 500, nil
	}

	if ok {
		response := model.ErrorResponse{Err: "Не удалось пройти капчу"}
		u.LogInfo(response)
		return model.User{}, 400, response
	}

	if response := u.ValidateSignUpData(newUser); response != nil {
		u.LogInfo(response)
		return model.User{}, 400, response
	}

	isSignedUp, response := u.IsAlreadySignedUp(newUser.Email)
	if response != nil && reflect.TypeOf(response) != reflect.TypeOf(model.ErrorDescriptionResponse{}) {
		u.LogInfo(response)
		return model.User{}, 500, nil
	}

	if isSignedUp {
		response = model.ErrorResponse{Err: "Вы уже зарегестрированы"}
		u.LogInfo(response)
		return model.User{}, 400, response
	}

	err = u.AddNewUser(&newUser)
	if err != nil {
		u.LogError(err)
		return model.User{}, 500, nil
	}

	newUser.Password = ""
	newUser.SecondPassword = ""
	newUser.PasswordHash = nil
	if len(newUser.Photos) == 0 {
		newUser.Photos = make([]uuid.UUID, 0)
	}

	return newUser, 200, nil
}

func (u *UserUsecase) CheckCaptch(token string) (bool, error) {
	url := "https://www.google.com/recaptcha/api/siteverify"

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return false, err
	}

	q := req.URL.Query()
	q.Add("secret", googleCaptchaSecret)
	q.Add("response", token)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	// unmarshall the response into a GoogleResponse
	var googleResponse model.GoogleCaptcha
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	err = json.Unmarshal(body, &googleResponse)
	if err != nil {
		return false, err
	}

	if !googleResponse.Success {
		return false, nil
	}

	return true, nil
}

func (u *UserUsecase) GetUserInfoById(id int) (user model.User, err error) {
	user, err = u.Db.GetUser(id)

	return user, err
}

func (u *UserUsecase) DeleteUser(id int) error {
	return u.Db.DeleteUser(id)
}

func (u *UserUsecase) SignIn(user model.User) (newUser model.User, code int, err error) {
	isValid, response := u.ValidateSignInData(user)
	if !isValid {
		u.LogInfo(response)
		return model.User{}, 400, response
	}

	isCorrect, err := u.CheckPasswordWithEmail(user.Password, user.Email)
	if err != nil {
		u.LogError(err)
		return model.User{}, 500, nil
	}
	if !isCorrect {
		response := model.ErrorResponse{Err: "Неверный логин или пароль"}
		u.LogInfo(response)
		return model.User{}, 401, response
	}

	newUser, err = u.Db.SignIn(user.Email)
	if err != nil {
		u.LogError(err)
		return model.User{}, 500, nil
	}

	if len(newUser.Photos) == 0 {
		newUser.Photos = make([]uuid.UUID, 0)
	}

	newUser.PasswordHash = nil

	return newUser, 200, nil
}

func (u *UserUsecase) ValidateSex(sex string) bool {
	if sex != "male" && sex != "female" {
		return false
	}

	return true
}

func (u *UserUsecase) ValidateDatePreferences(pref string) bool {
	if pref != "male" && pref != "female" && pref != "both" {
		return false
	}

	return true
}

func (u *UserUsecase) CheckPasswordWithId(passToCheck string, id int) (bool, error) {
	password, err := u.Db.GetPassWithId(id)
	if err != nil {
		return false, err
	}
	if password == nil {
		return false, nil
	}

	pass := sha3.New512()
	pass.Write([]byte(passToCheck))
	err = bcrypt.CompareHashAndPassword(password, pass.Sum(nil))
	if err != nil {
		return false, nil
	}

	return true, nil
}

func (u *UserUsecase) CheckPasswordWithEmail(passToCheck string, email string) (bool, error) {
	password, err := u.Db.GetPassWithEmail(email)
	if err != nil {
		return false, err
	}
	if password == nil {
		return false, nil
	}

	pass := sha3.New512()
	pass.Write([]byte(passToCheck))
	err = bcrypt.CompareHashAndPassword(password, pass.Sum(nil))
	if err != nil {
		return false, nil
	}

	return true, nil
}

func (u *UserUsecase) ChangeUserProperties(newUser *model.User) error {
	bufUser, err := u.Db.GetUser(newUser.Id)
	if err != nil {
		return err
	}

	if newUser.Email != "" {
		if !govalidator.IsEmail(newUser.Email) {
			response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Не удалось поменять данные"}
			response.Description["mail"] = "Почта не прошла валидацию"
			return response
		}
		isSignedUp, response := u.IsAlreadySignedUp(newUser.Email)
		if response != nil {
			return response
		}
		if isSignedUp {
			response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Не удалось поменять данные"}
			response.Description["mail"] = "Почта занята"
			return response
		}

		bufUser.Email = newUser.Email
	}

	if newUser.Name != "" {
		bufUser.Name = newUser.Name
	}

	if newUser.Birthday != 0 {
		bufUser.Birthday = newUser.Birthday
	}

	if newUser.Description != "" {
		bufUser.Description = newUser.Description
	}

	if newUser.City != "" {
		bufUser.City = newUser.City
	}

	if newUser.Instagram != "" {
		bufUser.Instagram = newUser.Instagram
	}

	response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Не удалось поменять данные"}
	if newUser.Sex != "" {
		if !u.ValidateSex(newUser.Sex) {
			response.Description["sex"] = "Неверно введен пол"
			return response
		}
		bufUser.Sex = newUser.Sex
	}

	if newUser.DatePreference != "" {
		if !u.ValidateDatePreferences(newUser.DatePreference) {
			response.Description["datePreferences"] = "Неверно введены предпочтения"
			return response
		}
		bufUser.DatePreference = newUser.DatePreference
	}

	err = u.IsActive(&bufUser)
	if err != nil {
		return err
	}

	err = u.Db.ChangeUser(bufUser)
	if err != nil {
		return err
	}

	*newUser = bufUser

	return nil
}

func (u *UserUsecase) ChangeUserPassword(newUser *model.User) error {
	response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}

	if !u.ValidatePassword(newUser.Password) {
		response.Description["password"] = "Введите пароль"
		return response
	}

	if newUser.SecondPassword != newUser.Password {
		response.Description["password"] = "Пароли не совпадают"
		return response
	}

	ok, err := u.CheckPasswordWithId(newUser.OldPassword, newUser.Id)
	if err != nil {
		return err
	}
	if !ok {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Отказано в доступе"}
		response.Description["password"] = "Неверный пароль"
		return response
	}

	hash, err := u.HashPassword(newUser.Password)
	if err != nil {
		return err
	}

	err = u.Db.ChangePassword(newUser.Id, hash)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserUsecase) ValidatePassword(password string) bool {
	if len(password) >= 8 && len(password) <= 64 {
		return true
	}
	return false
}

func (u *UserUsecase) ValidateSignInData(newUser model.User) (bool, error) {
	response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}

	_, err := govalidator.ValidateStruct(newUser)
	if err != nil {
		response.Description["mail"] = govalidator.ErrorByField(err, "email")
		response.Description["password"] = govalidator.ErrorByField(err, "password")
		return false, response
	}

	return true, nil
}

func (u *UserUsecase) ValidateSignUpData(newUser model.User) error {
	response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Не удалось зарегистрироваться"}

	_, err := govalidator.ValidateStruct(newUser)
	if err != nil {
		response.Description = govalidator.ErrorsByField(err)
		return response
	}

	if newUser.Password != newUser.SecondPassword {
		response.Description["password"] = "Пароли не совпадают"
		return response
	}

	return nil
}

func (u *UserUsecase) IsAlreadySignedUp(newEmail string) (bool, error) {
	isSignUp, err := u.Db.CheckMail(newEmail)

	return isSignUp, err
}

func (u *UserUsecase) HashPassword(pass string) ([]byte, error) {
	firstHash := sha3.New512()
	firstHash.Write([]byte(pass))
	secondHash, err := bcrypt.GenerateFromPassword(firstHash.Sum(nil), 14)
	if err != nil {
		return nil, err
	}

	return secondHash, nil
}

func (u *UserUsecase) IsActive(newUser *model.User) error {
	photos, err := u.Db.GetPhotos(newUser.Id)
	if err != nil {
		return err
	}

	if len(newUser.Name) != 0 && len(newUser.DatePreference) != 0 && len(newUser.Sex) != 0 && len(photos) != 0 {
		newUser.IsActive = true
		return nil
	}

	newUser.IsActive = false
	return nil
}

func (u *UserUsecase) AddNewUser(newUser *model.User) error {
	var err error
	newUser.PasswordHash, err = u.HashPassword(newUser.Password)
	if err != nil {
		return err
	}
	newUser.Password = ""
	newUser.SecondPassword = ""

	newUser.IsActive = false
	newUser.IsDeleted = false

	id, err := u.Db.AddUser(*newUser)
	if err != nil {
		return err
	}

	newUser.Id = id

	return nil
}

func (u *UserUsecase) SetCookie(token string) http.Cookie {
	return http.Cookie{
		Name:     "token",
		Value:    token,
		Domain:   "localhost:3000",
		Expires:  time.Now().AddDate(0, 0, 1),
		SameSite: http.SameSiteLaxMode,
	}
}

func (u *UserUsecase) ParseJsonToUser(body io.ReadCloser) (model.User, error) {
	var newUser model.User
	decoder := json.NewDecoder(body)
	err := decoder.Decode(&newUser)
	defer body.Close()

	newUser.Email = u.Sanitizer.Sanitize(newUser.Email)
	newUser.Name = u.Sanitizer.Sanitize(newUser.Name)
	newUser.City = u.Sanitizer.Sanitize(newUser.City)
	newUser.Instagram = u.Sanitizer.Sanitize(newUser.Instagram)
	newUser.Description = u.Sanitizer.Sanitize(newUser.Description)

	return newUser, err
}

func (u *UserUsecase) GetIdFromContext(ctx context.Context) (int, bool) {
	id, ok := ctx.Value(model.CtxUserId).(int)
	if !ok {
		return 0, false
	}
	return id, true
}
