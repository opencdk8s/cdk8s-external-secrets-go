package k8s


// LeaseList is a list of Lease objects.
// Experimental.
type KubeLeaseListProps struct {
	// Items is a list of schema objects.
	// Experimental.
	Items *[]*KubeLeaseProps `field:"required" json:"items" yaml:"items"`
	// Standard list metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// Experimental.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

