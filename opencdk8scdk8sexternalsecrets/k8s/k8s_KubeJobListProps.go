package k8s


// JobList is a collection of jobs.
// Experimental.
type KubeJobListProps struct {
	// items is the list of Jobs.
	// Experimental.
	Items *[]*KubeJobProps `field:"required" json:"items" yaml:"items"`
	// Standard list metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// Experimental.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

