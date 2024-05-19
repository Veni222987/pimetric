package counter

import (
	"fmt"
	"sync/atomic"

	"github.com/Veni222987/pimetric/api"
)

// 定义counter类型的metric数据结构

// Counter metric类型数据结构
type Counter struct {
	Name  string         `json:"name"`
	Help  string         `json:"help"`
	Type  api.MetricType `json:"type"`
	Value int64          `json:"value"`
}

// Help counter的help方法
func (c *Counter) GetHelp() string {
	return c.Help
}

// GetType counter的getType方法
func (c *Counter) GetType() api.MetricType {
	return api.MetricTypeCounter
}

// GetName counter的getName方法
func (c *Counter) GetName() string {
	return c.Name
}

// Incr counter的incr方法
func (c *Counter) Incr() error {
	atomic.AddInt64(&c.Value, 1)
	return nil
}

// Decr counter的decr方法
func (c *Counter) Decr() error {
	atomic.AddInt64(&c.Value, -1)
	return nil
}

func (c *Counter) GetValue() any {
	return c.Value
}

func (c *Counter) SetValue(value any) error {
	if v, ok := value.(int64); ok {
		c.Value = v
		return nil
	}
	return fmt.Errorf("value type error")
}
