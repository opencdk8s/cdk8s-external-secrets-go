package k8s


// StorageClassList is a collection of storage classes.
// Experimental.
type KubeStorageClassListProps struct {
	// Items is the list of StorageClasses.
	// Experimental.
	Items *[]*KubeStorageClassProps `field:"required" json:"items" yaml:"items"`
	// Standard list metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
	// Experimental.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

