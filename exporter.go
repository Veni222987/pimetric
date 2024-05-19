package pimetric

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
)

// GenHandler 指定appName生成handler，可以注册到已有的端口上
func GenHandler(appName string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		resMap := map[string]interface{}{
			appName: GetMetricsMap(),
		}
		resBytes, err := json.Marshal(resMap)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		w.Write(resBytes)
	}
}

// ExportMetrics 新建端口启动监听，返回端口
func ExportMetrics(appName string) string {
	http.HandleFunc("/metrics", GenHandler(appName))
	port := getRandomPort()
	for http.ListenAndServe(port, nil) != nil {
		port = getRandomPort()
	}
	return port
}

func getRandomPort() string {
	minPort := 32768
	maxPort := 65535
	// 生成一个32768-65536内的随机整数
	randomPort := rand.Intn(maxPort-minPort+1) + minPort
	return ":" + strconv.Itoa(randomPort)
}
