package models

import (
	"encoding/json"
	"net/http"
	"server/internal/pkg/utils/metrics"
)

type ErrorDescriptionResponse struct {
	Description map[string]string `json:"description"`
	Err         string            `json:"error"`
}

type ErrorResponse struct {
	Err string `json:"error"`
}

func (e ErrorResponse) Error() string {
	ret, _ := json.Marshal(e)

	return string(ret)
}

func (e ErrorDescriptionResponse) Error() string {
	ret, _ := json.Marshal(e)

	return string(ret)
}

func ResponseWithJson(w http.ResponseWriter, code int, body interface{}) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(body)
}

func ParseGrpcError(str string) ErrorResponse {
	var res ErrorResponse
	json.Unmarshal([]byte(str), &res)
	return res
}

type (
	logfunc      func()
	responsefunc func()
	metricfunc   func()
)

func LoggerFunc(body interface{}, logfunc func(string2 interface{})) logfunc {
	return func() {
		logfunc(body)
	}
}

func ResponseFunc(w http.ResponseWriter, code int, body interface{}) responsefunc {
	return func() {
		ResponseWithJson(w, code, body)
	}
}

func MetricFunc(code int, r *http.Request, err error) metricfunc {
	return func() {
		if code == 500 {
			metrics.CreateRequestErrors(r, err)
			return
		}
		metrics.CreateRequestHits(code, r)
	}
}

func Process(logfunc logfunc, responsefunc responsefunc, metricfunc metricfunc) {
	logfunc()
	responsefunc()
	metricfunc()
}
