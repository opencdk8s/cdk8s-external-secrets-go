package k8s


// PodsMetricSource indicates how to scale on a metric describing each pod in the current scale target (for example, transactions-processed-per-second).
//
// The values will be averaged together before being compared to the target value.
// Experimental.
type PodsMetricSourceV2Beta2 struct {
	// metric identifies the target metric by name and selector.
	// Experimental.
	Metric *MetricIdentifierV2Beta2 `field:"required" json:"metric" yaml:"metric"`
	// target specifies the target value for the given metric.
	// Experimental.
	Target *MetricTargetV2Beta2 `field:"required" json:"target" yaml:"target"`
}

