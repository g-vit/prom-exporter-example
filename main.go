package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Collector struct {
	firstMetric  *prometheus.Desc
	secondMetric *prometheus.Desc
}

func newCollector() *Collector {
	return &Collector{
		firstMetric: prometheus.NewDesc("first_metric",
			"First random metric",
			nil, nil,
		),
		secondMetric: prometheus.NewDesc("second_metric",
			"Second random metric",
			nil, nil,
		),
	}
}

func (collector *Collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.firstMetric
	ch <- collector.secondMetric
}

func (collector *Collector) Collect(ch chan<- prometheus.Metric) {
	var metricValue float64
	rand.Seed(time.Now().UnixNano())
	metricValue = rand.Float64() * 100

	m1 := prometheus.MustNewConstMetric(collector.firstMetric, prometheus.CounterValue, metricValue)
	m2 := prometheus.MustNewConstMetric(collector.secondMetric, prometheus.CounterValue, metricValue)
	ch <- m1
	ch <- m2
}

func main() {
	c := newCollector()
	prometheus.MustRegister(c)

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":3093", nil))
}
