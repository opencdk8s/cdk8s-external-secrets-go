package k8s


// SelfSubjectRulesReviewSpec defines the specification for SelfSubjectRulesReview.
// Experimental.
type SelfSubjectRulesReviewSpec struct {
	// Namespace to evaluate rules for.
	//
	// Required.
	// Experimental.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

