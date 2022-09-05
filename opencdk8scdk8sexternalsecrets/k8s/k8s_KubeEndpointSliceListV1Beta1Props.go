package k8s


// EndpointSliceList represents a list of endpoint slices.
// Experimental.
type KubeEndpointSliceListV1Beta1Props struct {
	// List of endpoint slices.
	// Experimental.
	Items *[]*KubeEndpointSliceV1Beta1Props `field:"required" json:"items" yaml:"items"`
	// Standard list metadata.
	// Experimental.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

