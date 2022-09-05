package k8s


// ScaleSpec describes the attributes of a scale subresource.
// Experimental.
type ScaleSpec struct {
	// desired number of instances for the scaled object.
	// Experimental.
	Replicas *float64 `field:"optional" json:"replicas" yaml:"replicas"`
}

