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
	// Incr increments the value of the metric.
	Incr() error
	// Decr decrements the value of the metric.
	Decr() error
	// GetValue returns the value of the metric.
	GetValue() any
	// SetValue sets the value of the metric.
	SetValue(value any) error
}
