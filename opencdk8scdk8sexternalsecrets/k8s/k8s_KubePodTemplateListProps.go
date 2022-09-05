package k8s


// PodTemplateList is a list of PodTemplates.
// Experimental.
type KubePodTemplateListProps struct {
	// List of pod templates.
	// Experimental.
	Items *[]*KubePodTemplateProps `field:"required" json:"items" yaml:"items"`
	// Standard list metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	// Experimental.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

