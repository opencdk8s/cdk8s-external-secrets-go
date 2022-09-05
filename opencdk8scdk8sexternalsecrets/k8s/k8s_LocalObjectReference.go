package k8s


// LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.
// Experimental.
type LocalObjectReference struct {
	// Name of the referent.
	//
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

