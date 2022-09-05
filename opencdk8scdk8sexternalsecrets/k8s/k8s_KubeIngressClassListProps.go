package k8s


// IngressClassList is a collection of IngressClasses.
// Experimental.
type KubeIngressClassListProps struct {
	// Items is the list of IngressClasses.
	// Experimental.
	Items *[]*KubeIngressClassProps `field:"required" json:"items" yaml:"items"`
	// Standard list metadata.
	// Experimental.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

