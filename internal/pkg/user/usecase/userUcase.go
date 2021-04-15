package usecase

import (
	"encoding/json"
	"io"
	"reflect"
	model "server/internal/pkg/models"
	"server/internal/pkg/user/repository"

	"github.com/microcosm-cc/bluemonday"

	"github.com/asaskevich/govalidator"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/sha3"
)

type UserUsecaseInterface interface {
	ValidateSex(sex string) bool
	ValidateDatePreferences(pref string) bool
	CheckPasswordWithId(passToCheck string, id int) (bool, error)
	CheckPasswordWithEmail(passToCheck string, email string) (bool, error)
	ChangeUserProperties(newUser *model.User) error
	ChangeUserPassword(newUser *model.User) error
	ValidatePassword(password string) bool
	ValidateSignInData(newUser model.User) (bool, error)
	ValidateSignUpData(newUser model.User) error
	IsAlreadySignedUp(newEmail string) (bool, error)
	HashPassword(pass string) ([]byte, error)
	isActive(newUser *model.User) error
	AddNewUser(newUser *model.User) error
	ParseJsonToUser(body io.ReadCloser) (model.User, error)

	LogInfo(body interface{})
	LogError(body interface{})

	Login(newUser *model.User) (code int, body interface{})
	SignUp(newUser *model.User) (code int, body interface{})
	UserInfo(id int) (model.User, error)
	DeleteUser(id int) error
	CreateFeed(id int, limitInt int) (code int, body interface{})
	AddPhoto(id int, image string) (int, error)
	GetPhoto(id int) (string, error)
	ChangeUserInfo(newUser *model.User) (code int, body interface{})
}

type UserUsecase struct {
	Clients   *map[int]*websocket.Conn
	Db        repository.UserRepositoryInterface
	Logger    *logrus.Entry
	Sanitizer *bluemonday.Policy
}

func (u *UserUsecase) ChangeUserInfo(newUser *model.User) (code int, body interface{}) {
	if newUser.Password != "" {
		err := u.ChangeUserPassword(newUser)
		if err != nil {
			return 400, err
		}
		newUser.Password = ""
		newUser.OldPassword = ""
		newUser.SecondPassword = ""
	}

	err := u.ChangeUserProperties(newUser)
	if err != nil {
		if reflect.TypeOf(err) != reflect.TypeOf(model.ErrorDescriptionResponse{}) {
			u.LogError(err)
			return 500, nil
		}
		return 400, nil
	}

	newUser.PasswordHash = nil
	return 200, newUser
}

func (u *UserUsecase) GetPhoto(id int) (string, error) {
	return u.Db.GetPhoto(id)
}

func (u *UserUsecase) AddPhoto(id int, image string) (int, error) {
	return u.Db.AddPhoto(id, image)
}

func (u *UserUsecase) UserInfo(id int) (model.User, error) {
	userInfo, err := u.Db.GetUser(id)
	return userInfo, err
}

func (u *UserUsecase) DeleteUser(id int) error {
	return u.Db.DeleteUser(id)
}

func (u *UserUsecase) CreateFeed(id int, limitInt int) (code int, body interface{}) {
	feed, err := u.Db.GetFeed(id, limitInt)
	if err != nil {
		u.LogError(err)
		return 500, nil
	}

	if len(feed) < limitInt {
		err = u.Db.CreateFeed(id)
		if err != nil {
			u.LogError(err)
			return 500, nil
		}

		feed, err = u.Db.GetFeed(id, limitInt)
		if err != nil {
			u.LogError(err)
			return 500, nil
		}
	}

	if len(feed) == 0 {
		err := u.Db.ClearFeed(id)
		if err != nil {
			u.LogError(err)
			return 500, nil
		}

		feed, err = u.Db.GetFeed(id, limitInt)
		if err != nil {
			u.LogError(err)
			return 500, nil
		}
	}

	if len(feed) == 0 {
		feed = make([]int, 0)
	}
	return 200, feed
}

func (u *UserUsecase) Login(newUser *model.User) (code int, body interface{}) {
	isValid, response := u.ValidateSignInData(*newUser)
	if !isValid {
		u.LogError(response)
		return 400, response
	}

	isCorrect, err := u.CheckPasswordWithEmail(newUser.Password, newUser.Email)
	if err != nil {
		u.LogError(response)
		return 500, nil
	}

	if !isCorrect {
		response := model.ErrorResponse{Err: "Неверный логин или пароль"}
		u.LogError(err)
		return 401, response
	}

	*newUser, err = u.Db.SignIn(newUser.Email)
	if err != nil {
		u.LogError(response)
		return 500, nil
	}

	if len(newUser.Photos) == 0 {
		newUser.Photos = make([]int, 0)
	}

	newUser.PasswordHash = nil
	return 200, newUser
}

func (u *UserUsecase) SignUp(newUser *model.User) (code int, body interface{}) {
	if response := u.ValidateSignUpData(*newUser); response != nil {
		u.LogInfo(model.UserErrorInvalidData)
		return 400, response
	}

	isSignedUp, response := u.IsAlreadySignedUp(newUser.Email)
	if response != nil && reflect.TypeOf(response) != reflect.TypeOf(model.ErrorDescriptionResponse{}) {
		u.LogError(response)
		return 500, nil
	}

	if isSignedUp {
		u.LogError("Already Signed Up")
		return 400, response
	}

	err := u.AddNewUser(newUser)
	if err != nil {
		u.LogError(err)
		return 500, nil
	}

	newUser.Password = ""
	newUser.SecondPassword = ""
	newUser.PasswordHash = nil
	if len(newUser.Photos) == 0 {
		newUser.Photos = make([]int, 0)
	}

	return 200, newUser
}

func (u UserUsecase) LogInfo(body interface{}) {
	u.Logger.Info(body)
}

func (u UserUsecase) LogError(body interface{}) {
	u.Logger.Error(body)
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

	err = u.isActive(&bufUser)
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

func (u *UserUsecase) IsAlreadySignedUp(newEmail string) (bool, error) {
	isSignUp, err := u.Db.CheckMail(newEmail)
	if err != nil {
		return true, err
	}
	if isSignUp == true {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Не удалось зарегистрироваться"}
		response.Description["mail"] = "Почта уже зарегистрирована"
		return true, response
	}

	return false, nil
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

func (u *UserUsecase) isActive(newUser *model.User) error {
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

	err = u.isActive(newUser)
	if err != nil {
		return err
	}

	id, err := u.Db.AddUser(*newUser)
	if err != nil {
		return err
	}

	newUser.Id = id

	return nil
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
