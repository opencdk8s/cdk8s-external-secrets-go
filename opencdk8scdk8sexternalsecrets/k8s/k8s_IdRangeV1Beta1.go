package k8s


// IDRange provides a min/max of an allowed range of IDs.
// Experimental.
type IdRangeV1Beta1 struct {
	// max is the end of the range, inclusive.
	// Experimental.
	Max *float64 `field:"required" json:"max" yaml:"max"`
	// min is the start of the range, inclusive.
	// Experimental.
	Min *float64 `field:"required" json:"min" yaml:"min"`
}

