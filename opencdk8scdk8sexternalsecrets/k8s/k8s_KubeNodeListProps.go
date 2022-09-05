package k8s


// NodeList is the whole list of all Nodes which have been registered with master.
// Experimental.
type KubeNodeListProps struct {
	// List of nodes.
	// Experimental.
	Items *[]*KubeNodeProps `field:"required" json:"items" yaml:"items"`
	// Standard list metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	// Experimental.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

