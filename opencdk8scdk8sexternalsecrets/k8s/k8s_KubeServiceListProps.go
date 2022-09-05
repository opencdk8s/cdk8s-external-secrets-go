package k8s


// ServiceList holds a list of services.
// Experimental.
type KubeServiceListProps struct {
	// List of services.
	// Experimental.
	Items *[]*KubeServiceProps `field:"required" json:"items" yaml:"items"`
	// Standard list metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	// Experimental.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

