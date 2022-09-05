package k8s


// PriorityClassList is a collection of priority classes.
// Experimental.
type KubePriorityClassListV1Alpha1Props struct {
	// items is the list of PriorityClasses.
	// Experimental.
	Items *[]*KubePriorityClassV1Alpha1Props `field:"required" json:"items" yaml:"items"`
	// Standard list metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
	// Experimental.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

