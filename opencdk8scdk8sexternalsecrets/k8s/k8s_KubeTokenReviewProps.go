package k8s


// TokenReview attempts to authenticate a token to a known user.
//
// Note: TokenReview requests may be cached by the webhook token authenticator plugin in the kube-apiserver.
// Experimental.
type KubeTokenReviewProps struct {
	// Spec holds information about the request being evaluated.
	// Experimental.
	Spec *TokenReviewSpec `field:"required" json:"spec" yaml:"spec"`
	// Standard object's metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// Experimental.
	Metadata *ObjectMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

