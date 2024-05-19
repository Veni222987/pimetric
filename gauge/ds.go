package gauge

import (
	"fmt"

	"github.com/Veni222987/pimetric/api"
)

// Gauge metric类型数据结构
type Gauge struct {
	Help  string         `json:"help"`
	Name  string         `json:"name"`
	Value float64        `json:"value"`
	Type  api.MetricType `json:"type"`
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

// Incr 增加指标的值(Gauge不支持)
func (g *Gauge) Incr() error {
	return fmt.Errorf("gauge can not incr")
}

// Decr 减少指标的值(Gauge不支持)
func (g *Gauge) Decr() error {
	return fmt.Errorf("gauge can not decr")
}

// GetValue 获取指标的值
func (g *Gauge) GetValue() any {
	return g.Value
}

// SetValue 设置指标的值
func (g *Gauge) SetValue(value any) error {
	if v, ok := value.(float64); ok {
		g.Value = v
		return nil
	}
	return fmt.Errorf("value type error")
}
