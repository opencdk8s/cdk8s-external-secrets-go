package k8s


// PodSecurityPolicyList is a list of PodSecurityPolicy objects.
// Experimental.
type KubePodSecurityPolicyListV1Beta1Props struct {
	// items is a list of schema objects.
	// Experimental.
	Items *[]*KubePodSecurityPolicyV1Beta1Props `field:"required" json:"items" yaml:"items"`
	// Standard list metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// Experimental.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

