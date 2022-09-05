package k8s


// LimitRange sets resource usage limits for each kind of resource in a Namespace.
// Experimental.
type KubeLimitRangeProps struct {
	// Standard object's metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// Experimental.
	Metadata *ObjectMeta `field:"optional" json:"metadata" yaml:"metadata"`
	// Spec defines the limits enforced.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	// Experimental.
	Spec *LimitRangeSpec `field:"optional" json:"spec" yaml:"spec"`
}

