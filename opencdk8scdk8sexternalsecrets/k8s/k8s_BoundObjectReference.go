package k8s


// BoundObjectReference is a reference to an object that a token is bound to.
// Experimental.
type BoundObjectReference struct {
	// API version of the referent.
	// Experimental.
	ApiVersion *string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
	// Kind of the referent.
	//
	// Valid kinds are 'Pod' and 'Secret'.
	// Experimental.
	Kind *string `field:"optional" json:"kind" yaml:"kind"`
	// Name of the referent.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// UID of the referent.
	// Experimental.
	Uid *string `field:"optional" json:"uid" yaml:"uid"`
}

