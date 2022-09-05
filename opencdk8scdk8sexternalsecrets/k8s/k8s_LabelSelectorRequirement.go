package k8s


// A label selector requirement is a selector that contains values, a key, and an operator that relates the key and values.
// Experimental.
type LabelSelectorRequirement struct {
	// key is the label key that the selector applies to.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// operator represents a key's relationship to a set of values.
	//
	// Valid operators are In, NotIn, Exists and DoesNotExist.
	// Experimental.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// values is an array of string values.
	//
	// If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.
	// Experimental.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

