package k8s

import (
	"time"
)

// The node this Taint is attached to has the "effect" on any pod that does not tolerate the Taint.
// Experimental.
type Taint struct {
	// Required.
	//
	// The effect of the taint on pods that do not tolerate the taint. Valid effects are NoSchedule, PreferNoSchedule and NoExecute.
	// Experimental.
	Effect *string `field:"required" json:"effect" yaml:"effect"`
	// Required.
	//
	// The taint key to be applied to a node.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// TimeAdded represents the time at which the taint was added.
	//
	// It is only written for NoExecute taints.
	// Experimental.
	TimeAdded *time.Time `field:"optional" json:"timeAdded" yaml:"timeAdded"`
	// The taint value corresponding to the taint key.
	// Experimental.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

