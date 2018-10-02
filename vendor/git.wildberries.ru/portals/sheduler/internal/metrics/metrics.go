package metrics

import (
	"strconv"
	"sync/atomic"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	// prometheusRegisterHTTPRequests - Количество HTTP запросов в секунду отновляется каждые <prometheusRegisterHTTPRequestsInterval> сек
	prometheusRegisterHTTPRequests prometheus.Gauge
	// PrometheusRegisterHTTPRequests - Количество HTTP запросов в секунду (текущее)
	PrometheusRegisterHTTPRequests int64
	// prometheusRegisterHTTPRequestsInterval - Интервал обновления статистики по HTTP запросам
	prometheusRegisterHTTPRequestsInterval = 10
)

// InitPrometheus Инициализация коллекторов метрик Prometheus
func InitPrometheus() {
	prometheusRegisterHTTPRequests = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "sheduler_go",
			Name:      "http_requests_per_second",
			Help:      "Количество HTTP запросов в секунду. Обновляется каждые " + strconv.Itoa(prometheusRegisterHTTPRequestsInterval) + " сек.",
		})
	prometheus.MustRegister(prometheusRegisterHTTPRequests)
	prometheusRegisterHTTPRequests.Set(0)

	countRequests()
}

func countRequests() {
	go func() {
		ticker := time.NewTicker(time.Duration(prometheusRegisterHTTPRequestsInterval) * time.Second)
		for range ticker.C {
			prometheusRegisterHTTPRequests.Set(float64(atomic.LoadInt64(&PrometheusRegisterHTTPRequests)) / float64(prometheusRegisterHTTPRequestsInterval))
			atomic.StoreInt64(&PrometheusRegisterHTTPRequests, 0)
		}
	}()
}
