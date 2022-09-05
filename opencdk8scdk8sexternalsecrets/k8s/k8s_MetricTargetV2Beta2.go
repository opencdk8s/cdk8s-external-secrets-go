package k8s


// MetricTarget defines the target value, average value, or average utilization of a specific metric.
// Experimental.
type MetricTargetV2Beta2 struct {
	// type represents whether the metric type is Utilization, Value, or AverageValue.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// averageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods.
	//
	// Currently only valid for Resource metric source type.
	// Experimental.
	AverageUtilization *float64 `field:"optional" json:"averageUtilization" yaml:"averageUtilization"`
	// averageValue is the target value of the average of the metric across all relevant pods (as a quantity).
	// Experimental.
	AverageValue Quantity `field:"optional" json:"averageValue" yaml:"averageValue"`
	// value is the target value of the metric (as a quantity).
	// Experimental.
	Value Quantity `field:"optional" json:"value" yaml:"value"`
}

