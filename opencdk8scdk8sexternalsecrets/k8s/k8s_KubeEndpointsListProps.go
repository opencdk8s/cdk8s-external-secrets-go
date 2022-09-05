package k8s


// EndpointsList is a list of endpoints.
// Experimental.
type KubeEndpointsListProps struct {
	// List of endpoints.
	// Experimental.
	Items *[]*KubeEndpointsProps `field:"required" json:"items" yaml:"items"`
	// Standard list metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	// Experimental.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

