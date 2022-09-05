package k8s


// PodsMetricSource indicates how to scale on a metric describing each pod in the current scale target (for example, transactions-processed-per-second).
//
// The values will be averaged together before being compared to the target value.
// Experimental.
type PodsMetricSourceV2Beta1 struct {
	// metricName is the name of the metric in question.
	// Experimental.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// targetAverageValue is the target value of the average of the metric across all relevant pods (as a quantity).
	// Experimental.
	TargetAverageValue Quantity `field:"required" json:"targetAverageValue" yaml:"targetAverageValue"`
	// selector is the string-encoded form of a standard kubernetes label selector for the given metric When set, it is passed as an additional parameter to the metrics server for more specific metrics scoping When unset, just the metricName will be used to gather metrics.
	// Experimental.
	Selector *LabelSelector `field:"optional" json:"selector" yaml:"selector"`
}

