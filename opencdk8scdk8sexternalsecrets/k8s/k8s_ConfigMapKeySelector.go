package k8s


// Selects a key from a ConfigMap.
// Experimental.
type ConfigMapKeySelector struct {
	// The key to select.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Name of the referent.
	//
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specify whether the ConfigMap or its key must be defined.
	// Experimental.
	Optional *bool `field:"optional" json:"optional" yaml:"optional"`
}

