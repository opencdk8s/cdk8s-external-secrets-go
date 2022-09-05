package k8s


// PriorityLevelConfigurationList is a list of PriorityLevelConfiguration objects.
// Experimental.
type KubePriorityLevelConfigurationListV1Beta1Props struct {
	// `items` is a list of request-priorities.
	// Experimental.
	Items *[]*KubePriorityLevelConfigurationV1Beta1Props `field:"required" json:"items" yaml:"items"`
	// `metadata` is the standard object's metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// Experimental.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

