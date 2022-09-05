package k8s


// PodTemplateSpec describes the data a pod should have when created from a template.
// Experimental.
type PodTemplateSpec struct {
	// Standard object's metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// Experimental.
	Metadata *ObjectMeta `field:"optional" json:"metadata" yaml:"metadata"`
	// Specification of the desired behavior of the pod.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	// Experimental.
	Spec *PodSpec `field:"optional" json:"spec" yaml:"spec"`
}

