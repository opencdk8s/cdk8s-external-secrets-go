package k8s


// NodeSpec describes the attributes that a node is created with.
// Experimental.
type NodeSpec struct {
	// Deprecated.
	//
	// If specified, the source of the node's configuration. The DynamicKubeletConfig feature gate must be enabled for the Kubelet to use this field. This field is deprecated as of 1.22: https://git.k8s.io/enhancements/keps/sig-node/281-dynamic-kubelet-configuration
	// Experimental.
	ConfigSource *NodeConfigSource `field:"optional" json:"configSource" yaml:"configSource"`
	// Deprecated.
	//
	// Not all kubelets will set this field. Remove field after 1.13. see: https://issues.k8s.io/61966
	// Experimental.
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// PodCIDR represents the pod IP range assigned to the node.
	// Experimental.
	PodCidr *string `field:"optional" json:"podCidr" yaml:"podCidr"`
	// podCIDRs represents the IP ranges assigned to the node for usage by Pods on that node.
	//
	// If this field is specified, the 0th entry must match the podCIDR field. It may contain at most 1 value for each of IPv4 and IPv6.
	// Experimental.
	PodCidRs *[]*string `field:"optional" json:"podCidRs" yaml:"podCidRs"`
	// ID of the node assigned by the cloud provider in the format: <ProviderName>://<ProviderSpecificNodeID>.
	// Experimental.
	ProviderId *string `field:"optional" json:"providerId" yaml:"providerId"`
	// If specified, the node's taints.
	// Experimental.
	Taints *[]*Taint `field:"optional" json:"taints" yaml:"taints"`
	// Unschedulable controls node schedulability of new pods.
	//
	// By default, node is schedulable. More info: https://kubernetes.io/docs/concepts/nodes/node/#manual-node-administration
	// Experimental.
	Unschedulable *bool `field:"optional" json:"unschedulable" yaml:"unschedulable"`
}

