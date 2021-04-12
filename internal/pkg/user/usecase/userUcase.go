package usecase

import (
	"encoding/json"
	"io"
	model "server/internal/pkg/models"
	"server/internal/pkg/user/repository"

	"github.com/asaskevich/govalidator"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/sha3"
)

type UserUsecase struct {
	Clients *map[int]*websocket.Conn
	Db      repository.UserRepository
	Logger  *logrus.Entry
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

func (u *UserUsecase) CheckPassword(newUser *model.User) (bool, error) {
	password, err := u.Db.GetPass(newUser.Email)
	if err != nil {
		return false, err
	}
	if password == nil {
		return false, nil
	}

	pass := sha3.New512()
	pass.Write([]byte(newUser.Password))
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

	isActive(&bufUser)

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

	ok, err := u.CheckPassword(newUser)
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

func isActive(newUser *model.User) {
	if len(newUser.Name) != 0 && len(newUser.DatePreference) != 0 && len(newUser.Sex) != 0 {
		newUser.IsActive = true
		return
	}

	newUser.IsActive = false
}

func (u *UserUsecase) AddNewUser(newUser *model.User) error {
	var err error
	newUser.PasswordHash, err = u.HashPassword(newUser.Password)
	if err != nil {
		return err
	}
	newUser.Password = ""
	newUser.SecondPassword = ""

	isActive(newUser)

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
	return newUser, err
}
