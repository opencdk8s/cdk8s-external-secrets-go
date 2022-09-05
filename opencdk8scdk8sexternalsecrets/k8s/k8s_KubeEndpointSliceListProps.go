package k8s


// EndpointSliceList represents a list of endpoint slices.
// Experimental.
type KubeEndpointSliceListProps struct {
	// List of endpoint slices.
	// Experimental.
	Items *[]*KubeEndpointSliceProps `field:"required" json:"items" yaml:"items"`
	// Standard list metadata.
	// Experimental.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

