package k8s


// DeploymentStrategy describes how to replace existing pods with new ones.
// Experimental.
type DeploymentStrategy struct {
	// Rolling update config params.
	//
	// Present only if DeploymentStrategyType = RollingUpdate.
	// Experimental.
	RollingUpdate *RollingUpdateDeployment `field:"optional" json:"rollingUpdate" yaml:"rollingUpdate"`
	// Type of deployment.
	//
	// Can be "Recreate" or "RollingUpdate". Default is RollingUpdate.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

