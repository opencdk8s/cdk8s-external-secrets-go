package k8s


// DaemonSetList is a collection of daemon sets.
// Experimental.
type KubeDaemonSetListProps struct {
	// A list of daemon sets.
	// Experimental.
	Items *[]*KubeDaemonSetProps `field:"required" json:"items" yaml:"items"`
	// Standard list metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// Experimental.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

