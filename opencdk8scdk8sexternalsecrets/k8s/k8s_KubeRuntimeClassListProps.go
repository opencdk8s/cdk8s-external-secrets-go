package k8s


// RuntimeClassList is a list of RuntimeClass objects.
// Experimental.
type KubeRuntimeClassListProps struct {
	// Items is a list of schema objects.
	// Experimental.
	Items *[]*KubeRuntimeClassProps `field:"required" json:"items" yaml:"items"`
	// Standard list metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// Experimental.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

