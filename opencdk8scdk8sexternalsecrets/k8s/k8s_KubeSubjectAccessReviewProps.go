package k8s


// SubjectAccessReview checks whether or not a user or group can perform an action.
// Experimental.
type KubeSubjectAccessReviewProps struct {
	// Spec holds information about the request being evaluated.
	// Experimental.
	Spec *SubjectAccessReviewSpec `field:"required" json:"spec" yaml:"spec"`
	// Standard list metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// Experimental.
	Metadata *ObjectMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

