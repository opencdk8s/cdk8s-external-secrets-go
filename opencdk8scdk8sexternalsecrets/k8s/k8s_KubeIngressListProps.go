package k8s


// IngressList is a collection of Ingress.
// Experimental.
type KubeIngressListProps struct {
	// Items is the list of Ingress.
	// Experimental.
	Items *[]*KubeIngressProps `field:"required" json:"items" yaml:"items"`
	// Standard object's metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// Experimental.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

