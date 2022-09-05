package k8s


// RuntimeClassSpec is a specification of a RuntimeClass.
//
// It contains parameters that are required to describe the RuntimeClass to the Container Runtime Interface (CRI) implementation, as well as any other components that need to understand how the pod will be run. The RuntimeClassSpec is immutable.
// Experimental.
type RuntimeClassSpecV1Alpha1 struct {
	// RuntimeHandler specifies the underlying runtime and configuration that the CRI implementation will use to handle pods of this class.
	//
	// The possible values are specific to the node & CRI configuration.  It is assumed that all handlers are available on every node, and handlers of the same name are equivalent on every node. For example, a handler called "runc" might specify that the runc OCI runtime (using native Linux containers) will be used to run the containers in a pod. The RuntimeHandler must be lowercase, conform to the DNS Label (RFC 1123) requirements, and is immutable.
	// Experimental.
	RuntimeHandler *string `field:"required" json:"runtimeHandler" yaml:"runtimeHandler"`
	// Overhead represents the resource overhead associated with running a pod for a given RuntimeClass.
	//
	// For more details, see https://git.k8s.io/enhancements/keps/sig-node/688-pod-overhead/README.md This field is beta-level as of Kubernetes v1.18, and is only honored by servers that enable the PodOverhead feature.
	// Experimental.
	Overhead *OverheadV1Alpha1 `field:"optional" json:"overhead" yaml:"overhead"`
	// Scheduling holds the scheduling constraints to ensure that pods running with this RuntimeClass are scheduled to nodes that support it.
	//
	// If scheduling is nil, this RuntimeClass is assumed to be supported by all nodes.
	// Experimental.
	Scheduling *SchedulingV1Alpha1 `field:"optional" json:"scheduling" yaml:"scheduling"`
}

