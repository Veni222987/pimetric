package historgram

import (
	"time"

	"github.com/Veni222987/pimetric/api"
)

type HistoPoint struct {
	Timestamp int64   `json:"time_stamp"`
	Number    float64 `json:"number"`
}

// Histogram metric类型数据结构
type Histogram struct {
	Name  string         `json:"name"`
	Help  string         `json:"help"`
	Type  api.MetricType `json:"type"`
	Value []HistoPoint   `json:"value"`
}

// GetHelp 获取metric的help信息
func (h *Histogram) GetHelp() string {
	return h.Help
}

// GetType 获取metric的类型
func (h *Histogram) GetType() api.MetricType {
	return h.Type
}

// GetName 获取metric的名称
func (h *Histogram) GetName() string {
	return h.Name
}

// GetValue 获取metric的值
func (h *Histogram) GetValue() any {
	return h.Value
}

// AddPoint 新增数据点
func (h *Histogram) AddPoint(number float64) {
	h.Value = append(h.Value, HistoPoint{
		Timestamp: time.Now().UnixMilli(),
		Number:    number,
	})
}
