package k8s


// A list of StorageVersions.
// Experimental.
type KubeStorageVersionListV1Alpha1Props struct {
	// Items holds a list of StorageVersion.
	// Experimental.
	Items *[]*KubeStorageVersionV1Alpha1Props `field:"required" json:"items" yaml:"items"`
	// Standard list metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// Experimental.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

