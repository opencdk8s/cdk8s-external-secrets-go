package k8s


// NetworkPolicy describes what network traffic is allowed for a set of Pods.
// Experimental.
type KubeNetworkPolicyProps struct {
	// Standard object's metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// Experimental.
	Metadata *ObjectMeta `field:"optional" json:"metadata" yaml:"metadata"`
	// Specification of the desired behavior for this NetworkPolicy.
	// Experimental.
	Spec *NetworkPolicySpec `field:"optional" json:"spec" yaml:"spec"`
}

