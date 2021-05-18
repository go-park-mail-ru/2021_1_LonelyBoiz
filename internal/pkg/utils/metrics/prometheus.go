package metrics

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	hits     *prometheus.CounterVec
	errors   *prometheus.CounterVec
	Duration *prometheus.HistogramVec
)

func New() {
	hits = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "hits",
	}, []string{"status", "path", "method"})

	errors = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "errors",
	}, []string{"path", "method", "error"})

	Duration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "request_duration",
		Buckets: []float64{0.001, 0.002, 0.005, 0.01, 0.02, 0.05, 0.1, 0.2, 0.5, 1, 2, 5, 10},
	}, []string{"URL", "method"})

	prometheus.MustRegister(hits, errors, Duration)
}

func CreateRequestHits(status int, r *http.Request) {
	hits.WithLabelValues(strconv.Itoa(status), strings.Split(r.URL.Path, "/")[1], r.Method).Inc()
}

func CreateRequestErrors(r *http.Request, err error) {
	errors.WithLabelValues(strings.Split(r.URL.Path, "/")[1], r.Method, err.Error()).Inc()
}
