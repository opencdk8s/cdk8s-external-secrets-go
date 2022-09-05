package k8s


// UserSubject holds detailed information for user-kind subject.
// Experimental.
type UserSubjectV1Beta1 struct {
	// `name` is the username that matches, or "*" to match all usernames.
	//
	// Required.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
}

