package k8s


// DeploymentList is a list of Deployments.
// Experimental.
type KubeDeploymentListProps struct {
	// Items is the list of Deployments.
	// Experimental.
	Items *[]*KubeDeploymentProps `field:"required" json:"items" yaml:"items"`
	// Standard list metadata.
	// Experimental.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

