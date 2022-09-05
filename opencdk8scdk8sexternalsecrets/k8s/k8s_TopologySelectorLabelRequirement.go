package k8s


// A topology selector requirement is a selector that matches given label.
//
// This is an alpha feature and may change in the future.
// Experimental.
type TopologySelectorLabelRequirement struct {
	// The label key that the selector applies to.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// An array of string values.
	//
	// One value must match the label to be selected. Each entry in Values is ORed.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

