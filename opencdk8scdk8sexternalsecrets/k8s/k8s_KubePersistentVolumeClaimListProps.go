package k8s


// PersistentVolumeClaimList is a list of PersistentVolumeClaim items.
// Experimental.
type KubePersistentVolumeClaimListProps struct {
	// A list of persistent volume claims.
	//
	// More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims
	// Experimental.
	Items *[]*KubePersistentVolumeClaimProps `field:"required" json:"items" yaml:"items"`
	// Standard list metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	// Experimental.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

