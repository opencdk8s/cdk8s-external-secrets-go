package k8s


// NetworkPolicyList is a list of NetworkPolicy objects.
// Experimental.
type KubeNetworkPolicyListProps struct {
	// Items is a list of schema objects.
	// Experimental.
	Items *[]*KubeNetworkPolicyProps `field:"required" json:"items" yaml:"items"`
	// Standard list metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// Experimental.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

