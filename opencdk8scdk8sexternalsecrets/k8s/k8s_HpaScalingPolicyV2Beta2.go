package k8s


// HPAScalingPolicy is a single policy which must hold true for a specified past interval.
// Experimental.
type HpaScalingPolicyV2Beta2 struct {
	// PeriodSeconds specifies the window of time for which the policy should hold true.
	//
	// PeriodSeconds must be greater than zero and less than or equal to 1800 (30 min).
	// Experimental.
	PeriodSeconds *float64 `field:"required" json:"periodSeconds" yaml:"periodSeconds"`
	// Type is used to specify the scaling policy.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Value contains the amount of change which is permitted by the policy.
	//
	// It must be greater than zero.
	// Experimental.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

