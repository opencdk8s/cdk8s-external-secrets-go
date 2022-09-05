package k8s


// ControllerRevisionList is a resource containing a list of ControllerRevision objects.
// Experimental.
type KubeControllerRevisionListProps struct {
	// Items is the list of ControllerRevisions.
	// Experimental.
	Items *[]*KubeControllerRevisionProps `field:"required" json:"items" yaml:"items"`
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
	// Experimental.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

