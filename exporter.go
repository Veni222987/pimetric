package pimetric

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/Veni222987/pimetric/counter"
	"github.com/Veni222987/pimetric/gauge"
	"github.com/Veni222987/pimetric/historgram"
	"github.com/Veni222987/pimetric/pimstore"
)

type Metricx struct {
	AppName      string                           `json:"app_name"`
	CounterMap   map[string]*counter.Counter      `json:"counter_map"`
	GaugeMap     map[string]*gauge.Gauge          `json:"gauge_map"`
	HistogramMap map[string]*historgram.Histogram `json:"histogram_map"`
}

// GenHandler 指定appName生成handler，可以注册到已有的端口上
func GenHandler(appName string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		resBytes, err := json.Marshal(Metricx{
			AppName:      appName,
			CounterMap:   pimstore.CounterMap,
			GaugeMap:     pimstore.GaugeMap,
			HistogramMap: pimstore.HistogramMap,
		})
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		pimstore.ClearHistogram()
		w.Write(resBytes)
	}
}

// ExportMetrics 新建端口启动监听
func ExportMetrics(appName string) {
	http.HandleFunc("/metrics", GenHandler(appName))
	err := http.ListenAndServe(":62888", nil)
	for err != nil {
		err = http.ListenAndServe(getRandomPort(), nil)
	}
}

func getRandomPort() string {
	minPort := 32768
	maxPort := 65535
	// 生成一个32768-65536内的随机整数
	randomPort := rand.Intn(maxPort-minPort+1) + minPort
	return ":" + strconv.Itoa(randomPort)
}
