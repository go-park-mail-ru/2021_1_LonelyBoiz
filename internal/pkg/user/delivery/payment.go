package delivery

import (
	"encoding/json"
	"net/http"
	"server/internal/pkg/models"
)

func (a *UserHandler) Payment(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		a.UserCase.LogError(err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	labelString := r.PostFormValue("label")
	if labelString == "" {
		a.UserCase.LogError("Пустой лэйбл")
		w.WriteHeader(400)
		return
	}

	var labelStruct models.Label
	err = json.Unmarshal([]byte(labelString), labelStruct)
	if err != nil {
		a.UserCase.LogError(err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	a.UserCase.LogInfo(labelStruct)

	err = a.UserCase.UpdatePayment(labelStruct.UserId, labelStruct.Amount)
	if err != nil {
		a.UserCase.LogError(err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}
