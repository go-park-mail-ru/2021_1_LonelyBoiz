package usecase

import (
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/sha3"
	"io"
	model "server/internal/pkg/models"
	"server/internal/pkg/user/repository"
)

type UserUsecase struct {
	Db     repository.UserRepository
	Logger *logrus.Entry
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

func (u *UserUsecase) checkPasswordForCHanging(newUser model.User) bool {
	oldUserPass, err := u.Db.GetPass(newUser.Id)
	if err != nil {
		return false
	}

	pass := sha3.New512()
	pass.Write([]byte(newUser.OldPassword))
	err = bcrypt.CompareHashAndPassword(oldUserPass, pass.Sum(nil))
	if err != nil {
		return false
	}

	return true
}

func (u *UserUsecase) ChangeUserProperties(newUser *model.User) error {
	bufUser, err := u.Db.GetUser(newUser.Id)
	if err != nil {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: err.Error()}
		response.Description["id"] = "Пользователя с таким id нет"
		return response
	}

	if newUser.Email != "" {
		if !govalidator.IsEmail(newUser.Email) {
			response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Отказано в доступе"}
			response.Description["mail"] = "Почта не прошла валидацию"
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

	if newUser.Avatar != "" {
		bufUser.Avatar = newUser.Avatar
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

	if len(newUser.PasswordHash) != 0 {
		bufUser.PasswordHash = newUser.PasswordHash
	}

	err = u.Db.ChangeUser(bufUser)
	if err != nil {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: err.Error()}
		return response
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

	if !u.checkPasswordForCHanging(*newUser) {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Отказано в доступе"}
		response.Description["password"] = "Неверный пароль"
		return response
	}

	hash, err := u.HashPassword(newUser.Password)
	if err != nil {
		response.Description["password"] = "Не удалось поменять пароль"
		return response
	}
	newUser.PasswordHash = hash

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

func (u *UserUsecase) CheckPassword(newUser *model.User) bool {
	user, err := u.Db.SignIn(newUser.Email)
	if err != nil || user.IsDeleted == true {
		return false
	}

	pass := sha3.New512()
	pass.Write([]byte(newUser.Password))
	err = bcrypt.CompareHashAndPassword(user.PasswordHash, pass.Sum(nil))
	if err != nil {
		return false
	}

	*newUser = user

	return true
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
	isSignUp := u.Db.CheckMail(newEmail)
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
	return secondHash, err
}

func (u *UserUsecase) AddNewUser(newUser *model.User) error {
	var err error
	newUser.PasswordHash, err = u.HashPassword(newUser.Password)
	if err != nil {
		return err
	}
	newUser.Password = ""
	newUser.SecondPassword = ""

	id, err := u.Db.AddUser(*newUser)
	if err != nil {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: err.Error()}
		return response
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
