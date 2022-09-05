package k8s


// DaemonSetUpdateStrategy is a struct used to control the update strategy for a DaemonSet.
// Experimental.
type DaemonSetUpdateStrategy struct {
	// Rolling update config params.
	//
	// Present only if type = "RollingUpdate".
	// Experimental.
	RollingUpdate *RollingUpdateDaemonSet `field:"optional" json:"rollingUpdate" yaml:"rollingUpdate"`
	// Type of daemon set update.
	//
	// Can be "RollingUpdate" or "OnDelete". Default is RollingUpdate.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

