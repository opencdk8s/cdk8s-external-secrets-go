package k8s


// ConfigMapList is a resource containing a list of ConfigMap objects.
// Experimental.
type KubeConfigMapListProps struct {
	// Items is the list of ConfigMaps.
	// Experimental.
	Items *[]*KubeConfigMapProps `field:"required" json:"items" yaml:"items"`
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
	// Experimental.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

