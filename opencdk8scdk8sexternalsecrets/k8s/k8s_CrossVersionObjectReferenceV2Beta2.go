package k8s


// CrossVersionObjectReference contains enough information to let you identify the referred resource.
// Experimental.
type CrossVersionObjectReferenceV2Beta2 struct {
	// Kind of the referent;
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds"
	// Experimental.
	Kind *string `field:"required" json:"kind" yaml:"kind"`
	// Name of the referent;
	//
	// More info: http://kubernetes.io/docs/user-guide/identifiers#names
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// API version of the referent.
	// Experimental.
	ApiVersion *string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
}

