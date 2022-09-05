package k8s


// ResourceQuotaList is a list of ResourceQuota items.
// Experimental.
type KubeResourceQuotaListProps struct {
	// Items is a list of ResourceQuota objects.
	//
	// More info: https://kubernetes.io/docs/concepts/policy/resource-quotas/
	// Experimental.
	Items *[]*KubeResourceQuotaProps `field:"required" json:"items" yaml:"items"`
	// Standard list metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	// Experimental.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

