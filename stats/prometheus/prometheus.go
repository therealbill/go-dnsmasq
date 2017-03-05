package stats

import "github.com/prometheus/client_golang/prometheus"

// prometheus metrics
var (
	RequestsGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "godnsmasq_requests",
		Help: "Number of requests handled by the server",
	})
	DNSSECRequestsGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "godnsmasq_requests_dnssec",
		Help: "Number of dnssec requests handled by the server",
	})
	CacheHitGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "godnsmasq_cache_hits",
		Help: "Number of cache hits",
	})
	CacheMissGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "godnsmasq_cache_misses",
		Help: "Number of cache misses",
	})
	RequestsDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "godnsmasq_request_duration_seconds",
		Help:    "Histogram of the request duration",
		Buckets: []float64{.01, .025, .05, .1, .5, 1.0, 2.0},
	})
	ForwardGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "godnsmasq_forwards",
		Help: "Number of forwards handled by the server",
	})
	StubForwardGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "godnsmasq_stub_forwards",
		Help: "Number of stub zone forwards handled by the server",
	})
)

func init() {
	prometheus.MustRegister(RequestsGauge)
	prometheus.MustRegister(DNSSECRequestsGauge)
	prometheus.MustRegister(CacheHitGauge)
	prometheus.MustRegister(CacheMissGauge)
	prometheus.MustRegister(RequestsDuration)
	prometheus.MustRegister(ForwardGauge)
}
