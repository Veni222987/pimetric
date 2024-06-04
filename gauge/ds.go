package gauge

import (
	"time"

	"github.com/Veni222987/pimetric/api"
)

// Gauge metric类型数据结构
type Gauge struct {
	Name      string         `json:"name"`
	Help      string         `json:"help"`
	Type      api.MetricType `json:"type"`
	Timestamp int64          `json:"time_stamp"`
	Value     float64        `json:"value"`
}

// GetHelp 获取指标的帮助信息
func (g *Gauge) GetHelp() string {
	return g.Help
}

// GetType 获取指标的类型
func (g *Gauge) GetType() api.MetricType {
	return g.Type
}

// GetName 获取指标的名称
func (g *Gauge) GetName() string {
	return g.Name
}

// GetTimestamp 获取时间戳的方法
func (c *Gauge) GetTimestamp() int64 {
	return c.Timestamp
}

// GetValue 获取指标的值
func (g *Gauge) GetValue() any {
	return g.Value
}

// SetValue 设置指标的值
func (g *Gauge) SetValue(value float64) error {
	g.Value = value
	g.Timestamp = time.Now().Unix()
	return nil
}
