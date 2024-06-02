package api

type MetricType string

const (
	MetricTypeCounter   MetricType = "counter"
	MetricTypeGauge     MetricType = "gauge"
	MetricTypeHistogram MetricType = "histogram"
	MetricTypeSummary   MetricType = "summary"
)

type Metric interface {
	// Help returns the help string for the metric.
	GetHelp() string
	// GetType returns the type of the metric.
	GetType() MetricType
	// GetName returns the name of the metric.
	GetName() string
	// GetTimestamp returns the timestamp of the metric
	GetTimestamp() int64
}
