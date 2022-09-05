package k8s


// SelfSubjectAccessReview checks whether or the current user can perform an action.
//
// Not filling in a spec.namespace means "in all namespaces".  Self is a special case, because users should always be able to check whether they can perform an action
// Experimental.
type KubeSelfSubjectAccessReviewProps struct {
	// Spec holds information about the request being evaluated.
	//
	// user and groups must be empty.
	// Experimental.
	Spec *SelfSubjectAccessReviewSpec `field:"required" json:"spec" yaml:"spec"`
	// Standard list metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// Experimental.
	Metadata *ObjectMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

