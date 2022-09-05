package k8s


// PodTemplate describes a template for creating copies of a predefined pod.
// Experimental.
type KubePodTemplateProps struct {
	// Standard object's metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// Experimental.
	Metadata *ObjectMeta `field:"optional" json:"metadata" yaml:"metadata"`
	// Template defines the pods that will be created from this pod template.
	//
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	// Experimental.
	Template *PodTemplateSpec `field:"optional" json:"template" yaml:"template"`
}

