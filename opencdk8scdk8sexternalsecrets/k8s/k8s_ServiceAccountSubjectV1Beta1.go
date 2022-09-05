package k8s


// ServiceAccountSubject holds detailed information for service-account-kind subject.
// Experimental.
type ServiceAccountSubjectV1Beta1 struct {
	// `name` is the name of matching ServiceAccount objects, or "*" to match regardless of name.
	//
	// Required.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `namespace` is the namespace of matching ServiceAccount objects.
	//
	// Required.
	// Experimental.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
}

