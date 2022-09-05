package k8s


// TypedLocalObjectReference contains enough information to let you locate the typed referenced object inside the same namespace.
// Experimental.
type TypedLocalObjectReference struct {
	// Kind is the type of resource being referenced.
	// Experimental.
	Kind *string `field:"required" json:"kind" yaml:"kind"`
	// Name is the name of resource being referenced.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// APIGroup is the group for the resource being referenced.
	//
	// If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.
	// Experimental.
	ApiGroup *string `field:"optional" json:"apiGroup" yaml:"apiGroup"`
}

