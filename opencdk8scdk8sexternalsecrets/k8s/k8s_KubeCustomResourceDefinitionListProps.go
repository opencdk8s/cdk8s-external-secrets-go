package k8s


// CustomResourceDefinitionList is a list of CustomResourceDefinition objects.
// Experimental.
type KubeCustomResourceDefinitionListProps struct {
	// items list individual CustomResourceDefinition objects.
	// Experimental.
	Items *[]*KubeCustomResourceDefinitionProps `field:"required" json:"items" yaml:"items"`
	// Standard object's metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
	// Experimental.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

