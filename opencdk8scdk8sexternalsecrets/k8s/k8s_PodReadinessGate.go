package k8s


// PodReadinessGate contains the reference to a pod condition.
// Experimental.
type PodReadinessGate struct {
	// ConditionType refers to a condition in the pod's condition list with matching type.
	// Experimental.
	ConditionType *string `field:"required" json:"conditionType" yaml:"conditionType"`
}

