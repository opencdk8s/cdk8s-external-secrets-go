package k8s


// LimitRangeSpec defines a min/max usage limit for resources that match on kind.
// Experimental.
type LimitRangeSpec struct {
	// Limits is the list of LimitRangeItem objects that are enforced.
	// Experimental.
	Limits *[]*LimitRangeItem `field:"required" json:"limits" yaml:"limits"`
}

