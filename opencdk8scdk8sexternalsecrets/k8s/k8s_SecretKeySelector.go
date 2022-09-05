package k8s


// SecretKeySelector selects a key of a Secret.
// Experimental.
type SecretKeySelector struct {
	// The key of the secret to select from.
	//
	// Must be a valid secret key.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Name of the referent.
	//
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specify whether the Secret or its key must be defined.
	// Experimental.
	Optional *bool `field:"optional" json:"optional" yaml:"optional"`
}

