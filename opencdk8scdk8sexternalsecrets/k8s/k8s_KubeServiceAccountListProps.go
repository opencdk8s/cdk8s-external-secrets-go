package k8s


// ServiceAccountList is a list of ServiceAccount objects.
// Experimental.
type KubeServiceAccountListProps struct {
	// List of ServiceAccounts.
	//
	// More info: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/
	// Experimental.
	Items *[]*KubeServiceAccountProps `field:"required" json:"items" yaml:"items"`
	// Standard list metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	// Experimental.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

