package k8s


// DeploymentSpec is the specification of the desired behavior of the Deployment.
// Experimental.
type DeploymentSpec struct {
	// Label selector for pods.
	//
	// Existing ReplicaSets whose pods are selected by this will be the ones affected by this deployment. It must match the pod template's labels.
	// Experimental.
	Selector *LabelSelector `field:"required" json:"selector" yaml:"selector"`
	// Template describes the pods that will be created.
	// Experimental.
	Template *PodTemplateSpec `field:"required" json:"template" yaml:"template"`
	// Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available.
	//
	// Defaults to 0 (pod will be considered available as soon as it is ready).
	// Experimental.
	MinReadySeconds *float64 `field:"optional" json:"minReadySeconds" yaml:"minReadySeconds"`
	// Indicates that the deployment is paused.
	// Experimental.
	Paused *bool `field:"optional" json:"paused" yaml:"paused"`
	// The maximum time in seconds for a deployment to make progress before it is considered to be failed.
	//
	// The deployment controller will continue to process failed deployments and a condition with a ProgressDeadlineExceeded reason will be surfaced in the deployment status. Note that progress will not be estimated during the time a deployment is paused. Defaults to 600s.
	// Experimental.
	ProgressDeadlineSeconds *float64 `field:"optional" json:"progressDeadlineSeconds" yaml:"progressDeadlineSeconds"`
	// Number of desired pods.
	//
	// This is a pointer to distinguish between explicit zero and not specified. Defaults to 1.
	// Experimental.
	Replicas *float64 `field:"optional" json:"replicas" yaml:"replicas"`
	// The number of old ReplicaSets to retain to allow rollback.
	//
	// This is a pointer to distinguish between explicit zero and not specified. Defaults to 10.
	// Experimental.
	RevisionHistoryLimit *float64 `field:"optional" json:"revisionHistoryLimit" yaml:"revisionHistoryLimit"`
	// The deployment strategy to use to replace existing pods with new ones.
	// Experimental.
	Strategy *DeploymentStrategy `field:"optional" json:"strategy" yaml:"strategy"`
}

