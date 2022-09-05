package k8s


// ConfigMapEnvSource selects a ConfigMap to populate the environment variables with.
//
// The contents of the target ConfigMap's Data field will represent the key-value pairs as environment variables.
// Experimental.
type ConfigMapEnvSource struct {
	// Name of the referent.
	//
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specify whether the ConfigMap must be defined.
	// Experimental.
	Optional *bool `field:"optional" json:"optional" yaml:"optional"`
}

