package k8s


// Sysctl defines a kernel parameter to be set.
// Experimental.
type Sysctl struct {
	// Name of a property to set.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Value of a property to set.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}

