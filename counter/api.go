package counter

var counterMap = make(map[string]int64)

// 获取统计数据
func GetCounter(metricName string) int64 {
	return counterMap[metricName]
}
