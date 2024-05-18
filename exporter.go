package main

import (
	"net/http"

	"github.com/Veni222987/pimetric/counter"
)

// ExportMetrics 启动监听
func ExportMetrics() {
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write(counter.GetSerializedCounters())
	})
	http.ListenAndServe(":62888", nil)
}
