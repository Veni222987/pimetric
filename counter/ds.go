package counter

import (
	"fmt"
	"sync/atomic"
	"time"

	"github.com/Veni222987/pimetric/api"
)

// 定义counter类型的metric数据结构

// Counter metric类型数据结构
type Counter struct {
	Name      string         `json:"name"`
	Help      string         `json:"help"`
	Type      api.MetricType `json:"type"`
	Timestamp int64          `json:"time_stamp"`
	Value     int64          `json:"value"`
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

// GetTimestamp 获取时间戳的方法
func (c *Counter) GetTimestamp() int64 {
	return c.Timestamp
}

// Incr counter的incr方法
func (c *Counter) Incr() error {
	atomic.AddInt64(&c.Value, 1)
	c.Timestamp = time.Now().Unix()
	return nil
}

// Decr counter的decr方法
func (c *Counter) Decr() error {
	atomic.AddInt64(&c.Value, -1)
	c.Timestamp = time.Now().Unix()
	return nil
}

// GetValue 获取counter值的方法
func (c *Counter) GetValue() any {
	return c.Value
}

// SetValue 设置counter值
func (c *Counter) SetValue(value any) error {
	if v, ok := value.(int64); ok {
		c.Value = v
		c.Timestamp = time.Now().Unix()
		return nil
	}
	return fmt.Errorf("value type error")
}
