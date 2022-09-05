package k8s


// Namespace provides a scope for Names.
//
// Use of multiple namespaces is optional.
// Experimental.
type KubeNamespaceProps struct {
	// Standard object's metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// Experimental.
	Metadata *ObjectMeta `field:"optional" json:"metadata" yaml:"metadata"`
	// Spec defines the behavior of the Namespace.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	// Experimental.
	Spec *NamespaceSpec `field:"optional" json:"spec" yaml:"spec"`
}

