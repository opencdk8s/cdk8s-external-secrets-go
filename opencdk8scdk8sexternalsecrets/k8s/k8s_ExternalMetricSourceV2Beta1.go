package k8s


// ExternalMetricSource indicates how to scale on a metric not associated with any Kubernetes object (for example length of queue in cloud messaging service, or QPS from loadbalancer running outside of cluster).
//
// Exactly one "target" type should be set.
// Experimental.
type ExternalMetricSourceV2Beta1 struct {
	// metricName is the name of the metric in question.
	// Experimental.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// metricSelector is used to identify a specific time series within a given metric.
	// Experimental.
	MetricSelector *LabelSelector `field:"optional" json:"metricSelector" yaml:"metricSelector"`
	// targetAverageValue is the target per-pod value of global metric (as a quantity).
	//
	// Mutually exclusive with TargetValue.
	// Experimental.
	TargetAverageValue Quantity `field:"optional" json:"targetAverageValue" yaml:"targetAverageValue"`
	// targetValue is the target value of the metric (as a quantity).
	//
	// Mutually exclusive with TargetAverageValue.
	// Experimental.
	TargetValue Quantity `field:"optional" json:"targetValue" yaml:"targetValue"`
}

