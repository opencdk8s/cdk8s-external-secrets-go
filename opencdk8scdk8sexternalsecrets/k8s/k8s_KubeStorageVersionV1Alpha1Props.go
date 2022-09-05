package k8s


// Storage version of a specific resource.
// Experimental.
type KubeStorageVersionV1Alpha1Props struct {
	// Spec is an empty spec.
	//
	// It is here to comply with Kubernetes API style.
	// Experimental.
	Spec interface{} `field:"required" json:"spec" yaml:"spec"`
	// The name is <group>.<resource>.
	// Experimental.
	Metadata *ObjectMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

