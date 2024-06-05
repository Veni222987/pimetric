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

var (
	CounterMap   map[string]*counter.Counter      = make(map[string]*counter.Counter)
	GaugeMap     map[string]*gauge.Gauge          = make(map[string]*gauge.Gauge)
	HistogramMap map[string]*historgram.Histogram = make(map[string]*historgram.Histogram)
)

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
	if value, ok := CounterMap[metricName]; ok {
		return value
	}
	rsp := &counter.Counter{
		Name:  metricName,
		Value: 0,
	}
	CounterMap[metricName] = rsp
	return rsp
}

// GaugeOf 获取gauge
func GaugeOf(metricName string) *gauge.Gauge {
	lock := lockMap[GAUGE_KEY]
	lock.Lock()
	defer lock.Unlock()
	if value, ok := GaugeMap[metricName]; ok {
		return value
	}
	rsp := &gauge.Gauge{
		Name:  metricName,
		Value: 0,
	}
	GaugeMap[metricName] = rsp
	return rsp
}

// HistogramOf 获取histogram
func HistogramOf(metricName string) *historgram.Histogram {
	lock := lockMap[HISTOGRAM_KEY]
	lock.Lock()
	defer lock.Unlock()
	if value, ok := HistogramMap[metricName]; ok {
		return value
	}
	rsp := &historgram.Histogram{
		Name:  metricName,
		Value: make([]historgram.HistoPoint, 0),
	}
	HistogramMap[metricName] = rsp
	return rsp
}

func ClearHistogram() {
	lock := lockMap[HISTOGRAM_KEY]
	lock.Lock()
	defer lock.Unlock()
	for _, v := range HistogramMap {
		v.Value = make([]historgram.HistoPoint, 0)
	}
}
