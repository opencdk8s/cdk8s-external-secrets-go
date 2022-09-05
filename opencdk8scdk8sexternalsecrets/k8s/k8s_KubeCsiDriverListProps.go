package k8s


// CSIDriverList is a collection of CSIDriver objects.
// Experimental.
type KubeCsiDriverListProps struct {
	// items is the list of CSIDriver.
	// Experimental.
	Items *[]*KubeCsiDriverProps `field:"required" json:"items" yaml:"items"`
	// Standard list metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
	// Experimental.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

