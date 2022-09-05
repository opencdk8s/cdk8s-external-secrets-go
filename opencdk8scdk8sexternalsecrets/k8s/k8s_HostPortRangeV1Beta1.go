package k8s


// HostPortRange defines a range of host ports that will be enabled by a policy for pods to use.
//
// It requires both the start and end to be defined.
// Experimental.
type HostPortRangeV1Beta1 struct {
	// max is the end of the range, inclusive.
	// Experimental.
	Max *float64 `field:"required" json:"max" yaml:"max"`
	// min is the start of the range, inclusive.
	// Experimental.
	Min *float64 `field:"required" json:"min" yaml:"min"`
}

