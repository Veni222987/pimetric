package counter

import (
	"encoding/json"
	"fmt"
	"log"
)

var counterMap = make(map[string]Counter)

// 我们的metric接口目前并不支持单独查询某一个指标的值，所以这里我们直接返回全部的counterMap
// GetSerializedCounters 返回json格式的counterMap
func GetSerializedCounters() []byte {
	jsonBytes, err := json.Marshal(counterMap)
	if err != nil {
		log.Printf("GetSerializedCounters error: %s\n", err)
		return []byte{}
	}
	return jsonBytes
}

// RegisterCounter 注册counter
func RegisterCounter(counter Counter) error {
	value, ok := counterMap[counter.Name]
	if !ok {
		return fmt.Errorf("counter name existed")
	}
	counterMap[counter.Name] = value
	return nil
}
