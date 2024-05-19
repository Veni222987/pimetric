package pimetric

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/Veni222987/pimetric/api"
	"github.com/Veni222987/pimetric/counter"
	"github.com/Veni222987/pimetric/gauge"
	"github.com/Veni222987/pimetric/historgram"
)

const (
	COUNTER_KEY   = "counter"
	GAUGE_KEY     = "gauge"
	HISTOGRAM_KEY = "histogram"
)

// metric不需要锁，初始化结束之后所有操作都是读指针，
var metricMap = make(map[string](map[string]api.Metric))

// 每一个类型的metric都需要一个锁
var lockMap = make(map[string]*sync.Mutex)

func init() {
	log.Println("Initializing pimetric")
	metricMap[COUNTER_KEY] = make(map[string]api.Metric)
	metricMap[GAUGE_KEY] = make(map[string]api.Metric)
	metricMap[HISTOGRAM_KEY] = make(map[string]api.Metric)

	lockMap[COUNTER_KEY] = &sync.Mutex{}
	lockMap[GAUGE_KEY] = &sync.Mutex{}
	lockMap[HISTOGRAM_KEY] = &sync.Mutex{}
}

// GetSerializedMetrics 返回json序列化字节流
func GetSerializedMetrics() []byte {
	var result []byte
	for _, v := range metricMap {
		for _, m := range v {
			tempByte, _ := json.Marshal(m)
			result = append(result, tempByte...)
		}
	}
	return result
}

// GetMetricsMap 返回metricMap
func GetMetricsMap() map[string](map[string]api.Metric) {
	lockMap[HISTOGRAM_KEY].Lock()
	defer lockMap[HISTOGRAM_KEY].Unlock()
	// 每次传出去之后historgram需要清空
	for _, v := range metricMap[HISTOGRAM_KEY] {
		if h, ok := v.(*historgram.Histogram); ok {
			h.Clear()
		}
	}
	return metricMap
}

// RegisterCounter 注册counter
func RegisterCounter(counter *counter.Counter) error {
	lock := lockMap[COUNTER_KEY]
	lock.Lock()
	defer lock.Unlock()
	metricMap[COUNTER_KEY][counter.GetName()] = counter
	return nil
}

// RegisterGauge 注册gauge
func RegisterGauge(gauge *gauge.Gauge) error {
	lock := lockMap[GAUGE_KEY]
	lock.Lock()
	defer lock.Unlock()
	metricMap[GAUGE_KEY][gauge.GetName()] = gauge
	return nil
}

// RegisterHistogram 注册histogram
func RegisterHistogram(histogram *historgram.Histogram) error {
	lock := lockMap[HISTOGRAM_KEY]
	lock.Lock()
	defer lock.Unlock()
	metricMap[HISTOGRAM_KEY][histogram.GetName()] = histogram
	return nil
}
