package k8s


// APIServiceList is a list of APIService objects.
// Experimental.
type KubeApiServiceListProps struct {
	// Items is the list of APIService.
	// Experimental.
	Items *[]*KubeApiServiceProps `field:"required" json:"items" yaml:"items"`
	// Standard list metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
	// Experimental.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

