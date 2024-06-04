package pimstore

import (
	"log"
	"sync"

	"github.com/Veni222987/pimetric/counter"
	"github.com/Veni222987/pimetric/gauge"
	"github.com/Veni222987/pimetric/historgram"
)

const (
	COUNTER_KEY   = "counter"
	GAUGE_KEY     = "gauge"
	HISTOGRAM_KEY = "histogram"
)

// MetricxStorage 所有metrics的结构体
type MetricxStorage struct {
	CounterMap   map[string]*counter.Counter      `json:"counter_map"`
	GaugeMap     map[string]*gauge.Gauge          `json:"gauge_map"`
	HistogramMap map[string]*historgram.Histogram `json:"histogram_map"`
}

var Metricx = &MetricxStorage{
	CounterMap:   make(map[string]*counter.Counter),
	GaugeMap:     make(map[string]*gauge.Gauge),
	HistogramMap: make(map[string]*historgram.Histogram),
}

// 每一个类型的metric都需要一个锁
var lockMap = make(map[string]*sync.Mutex)

func init() {
	log.Println("Initializing pimetric")
	lockMap[COUNTER_KEY] = &sync.Mutex{}
	lockMap[GAUGE_KEY] = &sync.Mutex{}
	lockMap[HISTOGRAM_KEY] = &sync.Mutex{}
}

// CounterOf 获取counter
func CounterOf(metricName string) *counter.Counter {
	lock := lockMap[COUNTER_KEY]
	lock.Lock()
	defer lock.Unlock()
	if value, ok := Metricx.CounterMap[metricName]; ok {
		return value
	}
	rsp := &counter.Counter{
		Name:  metricName,
		Value: 0,
	}
	Metricx.CounterMap[metricName] = rsp
	return rsp
}

// GaugeOf 获取gauge
func GaugeOf(metricName string) *gauge.Gauge {
	lock := lockMap[GAUGE_KEY]
	lock.Lock()
	defer lock.Unlock()
	if value, ok := Metricx.GaugeMap[metricName]; ok {
		return value
	}
	rsp := &gauge.Gauge{
		Name:  metricName,
		Value: 0,
	}
	Metricx.GaugeMap[metricName] = rsp
	return rsp
}

// HistogramOf 获取histogram
func HistogramOf(metricName string) *historgram.Histogram {
	lock := lockMap[HISTOGRAM_KEY]
	lock.Lock()
	defer lock.Unlock()
	if value, ok := Metricx.HistogramMap[metricName]; ok {
		return value
	}
	rsp := &historgram.Histogram{
		Name:  metricName,
		Value: make([]historgram.HistoPoint, 0),
	}
	Metricx.HistogramMap[metricName] = rsp
	return rsp
}
