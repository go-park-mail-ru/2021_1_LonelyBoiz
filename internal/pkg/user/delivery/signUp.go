package delivery

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"server/internal/pkg/models"
	model "server/internal/pkg/models"
)

func captchCheck(response string) (bool, error) {
	googleCaptchaSecret := "6LdIzK0aAAAAAPFpnVtQTL_r6gm7NnNhxZ5frVJq"
	url := "https://www.google.com/recaptcha/api/siteverify"

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return false, err
	}

	q := req.URL.Query()
	q.Add("secret", googleCaptchaSecret)
	q.Add("response", response)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	// unmarshall the response into a GoogleResponse
	var googleResponse models.GoogleCaptcha
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	err = json.Unmarshal(body, &googleResponse)
	if err != nil {
		return false, err
	}

	if googleResponse.Success == false {
		return false, nil
	}

	return true, nil
}

func (a *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	newUser, err := a.UserCase.ParseJsonToUser(r.Body)
	if err != nil {
		a.UserCase.Logger.Logger.Error(err)
		response := model.ErrorResponse{Err: "Не удалось прочитать тело запроса"}
		model.ResponseWithJson(w, 400, response)
		return
	}

	ok, err := captchCheck(newUser.CaptchaToken)
	if err != nil {
		a.UserCase.Logger.Logger.Error(err)
		model.ResponseWithJson(w, 500, nil)
		return
	}
	if !ok {
		response := model.ErrorResponse{Err: "Не удалось пройти капчу"}
		model.ResponseWithJson(w, 400, response)
		return
	}

	if response := a.UserCase.ValidateSignUpData(newUser); response != nil {
		model.ResponseWithJson(w, 400, response)
		a.UserCase.Logger.Info(model.UserErrorInvalidData)
		return
	}

	isSignedUp, response := a.UserCase.IsAlreadySignedUp(newUser.Email)
	if response != nil && reflect.TypeOf(response) != reflect.TypeOf(model.ErrorDescriptionResponse{}) {
		a.UserCase.Logger.Logger.Error(err)
		model.ResponseWithJson(w, 500, nil)
		return
	}
	if isSignedUp {
		a.UserCase.Logger.Info("Already Signed Up")
		model.ResponseWithJson(w, 400, response)
		return
	}

	err = a.UserCase.AddNewUser(&newUser)
	if err != nil {
		a.UserCase.Logger.Logger.Error(err)
		model.ResponseWithJson(w, 500, nil)
		return
	}

	err = a.Sessions.SetSession(w, newUser.Id)
	if err != nil {
		a.UserCase.Logger.Error(err)
		model.ResponseWithJson(w, 500, nil)
		return
	}

	newUser.Password = ""
	newUser.SecondPassword = ""
	newUser.PasswordHash = nil
	if len(newUser.Photos) == 0 {
		newUser.Photos = make([]int, 0)
	}
	model.ResponseWithJson(w, 200, newUser)
	a.UserCase.Logger.Info("Success SignUp")
}
