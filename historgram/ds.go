package historgram

import (
	"fmt"

	"github.com/Veni222987/pimetric/api"
)

// Histogram metric类型数据结构
type Histogram struct {
	Name  string         `json:"name"`
	Help  string         `json:"help"`
	Type  api.MetricType `json:"type"`
	Value []float64      `json:"value"`
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

// Incr 自增(histogram不支持)
func (h *Histogram) Incr() error {
	return fmt.Errorf("histogram can not incr")
}

// Decr 自减(histogram不支持)
func (h *Histogram) Decr() error {
	return fmt.Errorf("histogram can not decr")
}

// GetValue 获取metric的值
func (h *Histogram) GetValue() any {
	return h.Value
}

// SetValue 设置metric的值
func (h *Histogram) SetValue(value any) error {
	if v, ok := value.([]float64); ok {
		h.Value = v
		return nil
	}
	return fmt.Errorf("value type error")
}

func (h *Histogram) Clear() {
	h.Value = []float64{}
}
