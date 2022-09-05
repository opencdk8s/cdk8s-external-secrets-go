package k8s


// PriorityLevelConfigurationReference contains information that points to the "request-priority" being used.
// Experimental.
type PriorityLevelConfigurationReferenceV1Beta1 struct {
	// `name` is the name of the priority level configuration being referenced Required.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
}

