package counter

import "sync/atomic"

// 定义counter类型的metric数据结构

// Counter metric类型数据结构
type Counter struct {
	Name  string `json:"name"`
	Value int64  `json:"value"`
}

// Incr counter的incr方法
func (c *Counter) Incr() {
	atomic.AddInt64(&c.Value, 1)
}
