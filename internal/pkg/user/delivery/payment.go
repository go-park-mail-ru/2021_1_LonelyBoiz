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
		models.MetricFunc(400, r, err)
		return
	}

	labelString := r.PostFormValue("label")
	if labelString == "" {
		a.UserCase.LogError("Пустой лэйбл")
		w.WriteHeader(400)
		return
	}

	amountString := r.PostFormValue("withdraw_amount")
	if amountString == "" {
		a.UserCase.LogError("Пустой withdraw")
		w.WriteHeader(400)
		return
	}
	amountInt, ok := models.Tarif[amountString]
	if !ok {
		a.UserCase.LogError("Неправильный amount")
		w.WriteHeader(400)
		return
	}

	var labelStruct models.Label
	err = json.Unmarshal([]byte(labelString), &labelStruct)
	if err != nil {
		a.UserCase.LogError(err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		models.MetricFunc(400, r, err)
		return
	}

	err = a.UserCase.UpdatePayment(labelStruct.UserId, amountInt)
	if err != nil {
		a.UserCase.LogError(err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		models.MetricFunc(400, r, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	models.MetricFunc(200, r, nil)
	a.UserCase.LogInfo("success payment")
}
